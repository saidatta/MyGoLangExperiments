package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	client *http.Client
}

func NewUserService() *UserService {
	return &UserService{
		client: &http.Client{},
	}
}

func (s *UserService) GetUser(userID int) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://user-service-sidecar/users/%d", userID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

type OrderService struct {
	client *http.Client
}

func NewOrderService() *OrderService {
	return &OrderService{
		client: &http.Client{},
	}
}

func (s *OrderService) GetOrdersForUser(userID int) ([]*Order, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://order-service-sidecar/orders?userID=%d", userID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var orders []*Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, err
	}

	return orders, nil
}
