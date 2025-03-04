package main

import (
	"errors"                //used to simulate error
	"mocking-example/mocks" //Import the generated mock for Data Service
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProcessData_Success(t *testing.T) {
	ctrl := gomock.NewController(t) //Creates a new 'gomock' controller
	defer ctrl.Finish()             //Ensures all the expectations are met.

	mockService := mocks.NewMockDataService(ctrl) //Creates a mock instance of 'DataService'
	//Defines expected Bahvior: When GetData(1) is called, return "Mocked Data" and No error

	mockService.EXPECT().GetData(1).Return("Mocked Data", nil)

	//Create an instance of DataPrecessor with the mocked service
	processor := DataProcesssor{
		services: mockService,
	}

	//Call the function being tested.
	result, err := processor.ProcessDaata(1)

	//Verify the result.
	if err != nil || result != "Processed:Mocked Data" {
		t.Errorf("Expected 'Processed:Mocked Data', got %s", result)
	}

}
func TestProcessData_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDataService(ctrl)

	mockService.EXPECT().GetData(2).Return("", errors.New("Data not found"))

	processor := DataProcesssor{services: mockService}

	result, err := processor.ProcessDaata(2)

	if err == nil || err.Error() != "Data not found" {
		t.Errorf("Expected error 'Data not found',got %v", err)
	}

	if result != "" {
		t.Errorf("Expedted empty result, got %s", result)
	}
}

func TestProcessData_MultipleCalls(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDataService(ctrl)

	mockService.EXPECT().GetData(1).Return("Data1", nil).Times(1)

	mockService.EXPECT().GetData(2).Return("Data2", nil).Times(1)

	processor := DataProcesssor{
		services: mockService,
	}
	result1, _ := processor.ProcessDaata(1)
	result2, _ := processor.ProcessDaata(2)

	if result1 != "Processed:Data1" || result2 != "Processed:Data2" {
		t.Errorf("Unexpecred results: %s, %s", result1, result2)
	}

}
