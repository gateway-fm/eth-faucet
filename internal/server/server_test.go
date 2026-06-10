package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/negroni/v3"

	"github.com/chainflag/eth-faucet/internal/chain"
)

type MockTxBuilder struct {
	mock.Mock
}

func (m *MockTxBuilder) Sender() common.Address {
	args := m.Called()
	return args.Get(0).(common.Address)
}

func (m *MockTxBuilder) Transfer(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	args := m.Called(ctx, to, value)
	return args.Get(0).(common.Hash), args.Error(1)
}

func setupTestServer(mockBuilder chain.TxBuilder) *Server {
	cfg := &Config{
		httpPort:   8080,
		proxyCount: 0,
		interval:   0,
		network:    "testnet",
		symbol:     "ETH",
		payout:     1000000000, // 1 ETH
	}
	return NewServer(mockBuilder, cfg)
}

func TestHandleClaim(t *testing.T) {
	mockBuilder := new(MockTxBuilder)
	expectedAddress := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	expectedAmount := chain.EtherToWei(1.0)
	mockBuilder.On("Transfer", mock.Anything, expectedAddress, expectedAmount).Return(common.Hash{1}, nil)

	server := setupTestServer(mockBuilder)
	reqBody := strings.NewReader(fmt.Sprintf(`{"address": "%s"}`, expectedAddress))
	req, err := http.NewRequest("POST", "/api/claim", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := server.handleClaim()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, rr.Code)
	}

	var resp claimResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	mockBuilder.AssertExpectations(t)

}

// testAddress returns a valid EIP-55 checksummed address derived from n.
func testAddress(n int) string {
	return common.HexToAddress(fmt.Sprintf("0x%040x", n)).Hex()
}

// claim drives a single request through the limiter middleware as if it came
// from clientIP. nextStatus is the status the downstream handler responds with
// when the limiter lets the request through. It returns the final status code.
func claim(limiter *Limiter, address, clientIP string, nextStatus int) int {
	reqBody := strings.NewReader(fmt.Sprintf(`{"address": "%s"}`, address))
	req := httptest.NewRequest("POST", "/api/claim", reqBody)
	req.RemoteAddr = clientIP + ":12345"

	rr := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rr)
	limiter.ServeHTTP(rw, req, func(w http.ResponseWriter, r *http.Request) {
		renderJSON(w, claimResponse{Message: "ok"}, nextStatus)
	})
	return rw.Status()
}

func TestLimiterWalletAndIPWithdrawals(t *testing.T) {
	// 1 claim per wallet, 5 claims per IP, within an active window.
	limiter := NewLimiter(0, 5, 1, time.Hour)
	const ip = "10.0.0.1"

	// Five distinct wallets from the same IP all succeed.
	for i := 1; i <= 5; i++ {
		if code := claim(limiter, testAddress(i), ip, http.StatusOK); code != http.StatusOK {
			t.Fatalf("wallet %d: expected %d, got %d", i, http.StatusOK, code)
		}
	}

	// Sixth distinct wallet from the same IP is blocked by the IP limit.
	if code := claim(limiter, testAddress(6), ip, http.StatusOK); code != http.StatusTooManyRequests {
		t.Errorf("6th wallet: expected %d, got %d", http.StatusTooManyRequests, code)
	}

	// A fresh IP, but reusing wallet #1, is blocked by the per-wallet limit of 1.
	if code := claim(limiter, testAddress(1), "10.0.0.2", http.StatusOK); code != http.StatusTooManyRequests {
		t.Errorf("repeat wallet: expected %d, got %d", http.StatusTooManyRequests, code)
	}
}

func TestLimiterSingleWithdrawalRegression(t *testing.T) {
	// Defaults (1/1) preserve the original single-claim-per-window behavior.
	limiter := NewLimiter(0, 1, 1, time.Hour)
	const ip = "10.0.0.1"

	if code := claim(limiter, testAddress(1), ip, http.StatusOK); code != http.StatusOK {
		t.Fatalf("first claim: expected %d, got %d", http.StatusOK, code)
	}
	if code := claim(limiter, testAddress(2), ip, http.StatusOK); code != http.StatusTooManyRequests {
		t.Errorf("second claim same IP: expected %d, got %d", http.StatusTooManyRequests, code)
	}
}

func TestLimiterRollbackOnFailure(t *testing.T) {
	// A failed downstream response must not consume the window slot.
	limiter := NewLimiter(0, 1, 1, time.Hour)
	const ip = "10.0.0.1"
	addr := testAddress(1)

	if code := claim(limiter, addr, ip, http.StatusInternalServerError); code != http.StatusInternalServerError {
		t.Fatalf("failed claim: expected %d, got %d", http.StatusInternalServerError, code)
	}
	// The count was rolled back, so a retry is allowed.
	if code := claim(limiter, addr, ip, http.StatusOK); code != http.StatusOK {
		t.Errorf("retry after failure: expected %d, got %d", http.StatusOK, code)
	}
}

func TestHandleInfo(t *testing.T) {
	mockBuilder := new(MockTxBuilder)
	mockBuilder.On("Sender").Return(common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"))

	server := setupTestServer(mockBuilder)
	req, err := http.NewRequest("GET", "/api/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := server.handleInfo()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, rr.Code)
	}

	var resp infoResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	mockBuilder.AssertExpectations(t)
}
