package main

import (
	"encoding/json"
	"fmt"
	"os"

	employee "protobuf/go_pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	emp1 := &employee.Employee{
		Id:          1,
		Name:        "Fahim",
		Designation: "Software Engineer",
		Salary:      100.0,
	}
	// fmt.Println(emp1)
	// Marshal the employee object to binary format
	data, err := proto.Marshal(emp1)
	if err != nil {
		fmt.Println("Error marshalling employee:", err)
		return
	}
	fmt.Println("Marshalled data:", data)
	fmt.Println("Marshalled data length:", len(data))

	// Unmarshal the binary data back to an employee object
	emp2 := &employee.Employee{}
	err = proto.Unmarshal(data, emp2)
	if err != nil {
		fmt.Println("Error unmarshalling employee:", err)
		return
	}
	fmt.Println("Unmarshalled employee:", emp2)
	fmt.Println("Unmarshalled employee ID:", emp2.GetId())
	fmt.Println("Unmarshalled employee ID are same: ", emp2.Id == emp1.Id)
	fmt.Println("------------------------")
	fmt.Println("Writing employee to file...")
	if err := WriteEmployeeFiles(emp1, "employee.bin", "employee.json"); err != nil {
		fmt.Println("Error writing employee files:", err)
		return
	}
	fmt.Println("------------------------")
	fmt.Println("Employee files written successfully.")
}

func WriteEmployeeFiles(emp *employee.Employee, binFile, jsonFile string) error {
	// Marshal to protobuf binary
	data, err := proto.Marshal(emp)
	if err != nil {
		return fmt.Errorf("error marshalling employee: %w", err)
	}
	if err := os.WriteFile(binFile, data, 0644); err != nil {
		return fmt.Errorf("error writing binary file: %w", err)
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(emp)
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %w", err)
	}
	if err := os.WriteFile(jsonFile, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}

	// Check file sizes
	binInfo, err := os.Stat(binFile)
	if err != nil {
		return fmt.Errorf("error getting binary file info: %w", err)
	}
	jsonInfo, err := os.Stat(jsonFile)
	if err != nil {
		return fmt.Errorf("error getting JSON file info: %w", err)
	}

	fmt.Printf("Binary file size: %d bytes\n", binInfo.Size())
	fmt.Printf("JSON file size: %d bytes\n", jsonInfo.Size())
	return nil
}
