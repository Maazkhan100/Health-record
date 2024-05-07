/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a property
type SmartContract struct {
	contractapi.Contract
}

type Doctor struct {
	Doctor_id      string   `json:"doctor_id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	CNIC           string   `json:"cnic"`
	Password       string   `json:"password"`
	Phone          string   `json:"phone"`
	Hospital       string   `json:"hospital"`
	Role           string   `json:"role"`
	Specialization string   `json:"specialization"`
	Verified       bool     `json:"verified"`
	Patients       []string `json:"patients"`
	List           []string `json:"list"`
	Requests       []string `json:"requests"`
}

type Report struct {
	Report_id string `json:"report_id"`
	Doctor    string `json:"doctor"`
	Patient   string `json:"patient"`
	Date      string `json:"date"`
	Symptoms  string `json:"symptopms"`
	Advice    string `json:"advice"`
	Tests     []Test
	Vitals    string `json:"vitals"`
}

type Vitals struct {
	Vitals_id string `json:"vitals_id"`
	Patient   string `json:"patient"`
	Date      string `json:"date"`
	BloodPres string `json:"BloodPres"`
	Temp      string `json:"temp"`
	Weight    string `json:"weight"`
	Height    string `json:"height"`
}

type Appointment struct {
	Appointment_id string   `json:"Appointment_id"`
	Doctor         string   `json:"doctor"`
	Patient        string   `json:"patient"`
	Day            string   `json:"day"`
	Time           string   `json:"time"`
	Patients       []string `json:"patients"`
	List           []string `json:"list"`
	Requests       []string `json:"requests"`
}

type Nurse struct {
	Nurse_id string   `json:"nurse_id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	CNIC     string   `json:"cnic"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"`
	Hospital string   `json:"hospital"`
	Role     string   `json:"role"`
	Doctor   string   `json:"doctor"`
	Verified bool     `json:"verified"`
	Patients []string `json:"patients"`
	List     []string `json:"list"`
	Requests []string `json:"requests"`
}
type Pharmacist struct {
	Pharmacist_id string   `json:"pharmacist_id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	CNIC          string   `json:"cnic"`
	Password      string   `json:"password"`
	Phone         string   `json:"phone"`
	Hospital      string   `json:"hospital"`
	Role          string   `json:"role"`
	Verified      bool     `json:"verified"`
	Patients      []string `json:"patients"`
	List          []string `json:"list"`
	Requests      []string `json:"requests"`
}
type LabTechnician struct {
	Technician_id string   `json:"technician_id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	CNIC          string   `json:"cnic"`
	Password      string   `json:"password"`
	Phone         string   `json:"phone"`
	Hospital      string   `json:"hospital"`
	Role          string   `json:"role"`
	Verified      bool     `json:"verified"`
	Patients      []string `json:"patients"`
	List          []string `json:"list"`
	Requests      []string `json:"requests"`
}
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	CNIC     string `json:"cnic"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}
type Admin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Patient struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	CNIC     string `json:"cnic"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
	Age      string `json:"age"`
	Gender   string `json:"gender"`
	Weight   string `json:"weight"`
	Height   string `json:"height"`
	Drugs    []Drug
	Vitals   string   `json:"vitals"`
	List     []string `json:"list"`
	Reports  []string `json:"reports"`
	Requests []string `json:"requests"`
}

type Drug struct {
	Name     string `json:"name"`
	Timing   string `json:"timing"`
	Meal     string `json:"meal"`
	Duration string `json:"Duration"`
	Date     string `json:"Date"`
	Given    bool   `json:"given"`
}

type Test struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type QueryPatient struct {
	Key    string `json:"Key"`
	Record *Patient
}

type QueryAppointment struct {
	Key    string `json:"Key"`
	Record *Appointment
}

type QueryDoctor struct {
	Key    string `json:"Key"`
	Record *Doctor
}
type QueryNurse struct {
	Key    string `json:"Key"`
	Record *Nurse
}
type QueryTechnician struct {
	Key    string `json:"Key"`
	Record *LabTechnician
}
type QueryPharmacist struct {
	Key    string `json:"Key"`
	Record *Pharmacist
}
type QueryReport struct {
	Key    string `json:"Key"`
	Record *Report
}
type QueryVitals struct {
	Key    string `json:"Key"`
	Record *Vitals
}
type QueryDrug struct {
	Key    string `json:"Key"`
	Record *Drug
}

//InitLedger adds a base set of doctors and nurses to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	doctors := []Doctor{
		Doctor{Doctor_id: "24431", Name: "Maaz Khan", Email: "maaz@gmail.com", Role: "doctor", Verified: true, Patients: []string{"Saad", "Waleed"}, Requests: []string{""}, List: []string{""}, CNIC: "3374623854726", Phone: "03013358777", Hospital: "hospital1", Password: "1234", Specialization: "Physician"},
		Doctor{Doctor_id: "24432", Name: "Sara Jameel", Email: "Sara@gmail.com", Role: "doctor", Verified: true, Patients: []string{"Saad", "Waleed"}, Requests: []string{""}, List: []string{""}, CNIC: "3374623887452", Phone: "03013358777", Hospital: "hospital1", Password: "4321", Specialization: "Surgeon"},
	}

	nurses := []Nurse{
		Nurse{Nurse_id: "13331", Name: "Sadia Malik", Email: "sadia@gmail.com", Role: "nurse", Verified: true, Patients: []string{""}, Requests: []string{""}, List: []string{""}, CNIC: "3374623754698", Phone: "03013358777", Hospital: "hospital1", Password: "1234", Doctor: "Maaz Khan"},
		Nurse{Nurse_id: "13332", Name: "Maha Khan", Email: "maha@gmail.com", Role: "nurse", Verified: true, Patients: []string{""}, Requests: []string{""}, List: []string{""}, CNIC: "3374623887452", Phone: "03013358777", Hospital: "hospital1", Password: "4321", Doctor: "Sara Jameel"},
	}

	pharmacists := []Pharmacist{
		Pharmacist{Pharmacist_id: "34331", Name: "Laiba Jadoon", Email: "laiba@gmail.com", Role: "pharmacist", Verified: true, Patients: []string{""}, Requests: []string{""}, List: []string{""}, CNIC: "3374623774595", Phone: "03013358777", Hospital: "hospital1", Password: "1234"},
		Pharmacist{Pharmacist_id: "34332", Name: "Hamza Malik", Email: "hamza@gmail.com", Role: "pharmacist", Verified: true, Patients: []string{""}, Requests: []string{""}, List: []string{""}, CNIC: "3374623887444", Phone: "03013358777", Hospital: "hospital1", Password: "4321"},
	}

	technicians := []LabTechnician{
		LabTechnician{Technician_id: "45331", Name: "Fatima Khan", Email: "fatima@gmail.com", Role: "labtechnician", Verified: true, Patients: []string{""}, Requests: []string{""}, List: []string{""}, CNIC: "3374623774595", Phone: "03013358777", Hospital: "hospital1", Password: "1234"},
		LabTechnician{Technician_id: "45332", Name: "Ahmed Nazeem", Email: "ahmed@gmail.com", Role: "labtechnician", Verified: true, Patients: []string{""}, Requests: []string{""}, List: []string{""}, CNIC: "3374623887444", Phone: "03013358777", Hospital: "hospital1", Password: "4321"},
	}

	patients := []Patient{
		Patient{Name: "Imran Hameed", Email: "imran@gmail.com", CNIC: "3374623444595", Phone: "0301555555", Password: "patient1", Role: "patient", Address: "x-y-z", City: "Fasilabadh", Province: "Punjab", Age: "40", Gender: "MALE", Weight: "60kg", Height: "190cm", Drugs: []Drug{}, Vitals: "VITALS0", List: []string{""}, Requests: []string{""}, Reports: []string{""}},
		Patient{Name: "Aroona Sohail", Email: "aroona@gmail.com", CNIC: "3374623444596", Phone: "0301666666", Password: "patient2", Role: "patient", Address: "x-y-z", City: "Fasilabadh", Province: "Punjab", Age: "40", Gender: "FEMALE", Weight: "60kg", Height: "190cm", Photo: "./images/aroona.png", Drugs: []Drug{}, Vitals: "VITALS1", List: []string{""}, Requests: []string{""}, Reports: []string{"REPORT0", "REPORT1"}},
	}

	reports := []Report{
		Report{Report_id: "REPORT0", Doctor: "Sara Jameel", Patient: "Aroona Sohail", Date: "20/02/2022", Symptoms: "Sour Throat,Headaches", Advice: "", Tests: []Test{Test{Name: "CT Scan", Path: ""}, Test{Name: "Ultrasound", Path: ""}, Test{Name: "X-Ray", Path: ""}}, Vitals: "VITALS0"},
		Report{Report_id: "REPORT1", Doctor: "Sara Jameel", Patient: "Aroona Sohail", Date: "25/06/2022", Symptoms: "Sour Throat,Headaches", Advice: "", Tests: []Test{Test{Name: "Endoscopy", Path: ""}, Test{Name: "BloodTest", Path: ""}, Test{Name: "X-Ray", Path: ""}}, Vitals: "VITALS1"},
	}

	vitals := []Vitals{
		Vitals{Vitals_id: "VITALS0", Patient: "Aroona Sohail", Date: "20/02/2022", BloodPres: "122/70", Temp: "37", Weight: "47", Height: "170"},
		Vitals{Vitals_id: "VITALS1", Patient: "Aroona Sohail", Date: "25/06/2022", BloodPres: "121/75", Temp: "35", Weight: "50", Height: "170"},
	}

	for i, doctor := range doctors {
		doctorAsBytes, _ := json.Marshal(doctor)
		err := ctx.GetStub().PutState("DOCTOR"+strconv.Itoa(i), doctorAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	for i, nurse := range nurses {
		nurseAsBytes, _ := json.Marshal(nurse)
		err := ctx.GetStub().PutState("NURSE"+strconv.Itoa(i), nurseAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	for i, technician := range technicians {
		technicianAsBytes, _ := json.Marshal(technician)
		err := ctx.GetStub().PutState("LTECHNICIAN"+strconv.Itoa(i), technicianAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	for i, pharmacist := range pharmacists {
		pharmacistAsBytes, _ := json.Marshal(pharmacist)
		err := ctx.GetStub().PutState("PHARMACIST"+strconv.Itoa(i), pharmacistAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	for i, patient := range patients {
		patientAsBytes, _ := json.Marshal(patient)
		err := ctx.GetStub().PutState("PATIENT"+strconv.Itoa(i), patientAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	for i, report := range reports {
		reportAsBytes, _ := json.Marshal(report)
		err := ctx.GetStub().PutState("REPORT"+strconv.Itoa(i), reportAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	for i, vital := range vitals {
		vitalAsBytes, _ := json.Marshal(vital)
		err := ctx.GetStub().PutState("VITALS"+strconv.Itoa(i), vitalAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, userid string, name string, email string, cnic string, phone string, hospital string, password string, isAdmin bool) error {

	user := User{

		Name:     name,
		Email:    email,
		CNIC:     cnic,
		Password: password,
		IsAdmin:  isAdmin,
	}

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userid, userAsBytes)
}

func (s *SmartContract) CreateDoctor(ctx contractapi.TransactionContextInterface, userid string, name string, email string, cnic string, phone string, hospital string, password string, specialization string, id string) error {

	doctor := Doctor{
		Doctor_id:      userid,
		Name:           name,
		Email:          email,
		CNIC:           cnic,
		Phone:          phone,
		Hospital:       hospital,
		Password:       password,
		Specialization: specialization,
		Role:           "doctor",
		Verified:       false,
		Patients:       []string{""},
		List:           []string{""},
		Requests:       []string{""},
	}

	userAsBytes, _ := json.Marshal(doctor)

	return ctx.GetStub().PutState(userid, userAsBytes)
}

func (s *SmartContract) CreateNurse(ctx contractapi.TransactionContextInterface, userid string, name string, email string, cnic string, phone string, hospital string, password string, doctor string, id string) error {

	nurse := Nurse{
		Nurse_id: id,
		Name:     name,
		Email:    email,
		CNIC:     cnic,
		Phone:    phone,
		Hospital: hospital,
		Password: password,
		Doctor:   doctor,
		Role:     "nurse",
		Verified: false,
		Patients: []string{""},
		List:     []string{""},
		Requests: []string{""},
	}

	userAsBytes, _ := json.Marshal(nurse)

	return ctx.GetStub().PutState(userid, userAsBytes)
}

func (s *SmartContract) CreatePharmacist(ctx contractapi.TransactionContextInterface, userid string, name string, email string, cnic string, phone string, hospital string, password string, id string) error {

	pharmacist := Pharmacist{
		Pharmacist_id: id,
		Name:          name,
		Email:         email,
		CNIC:          cnic,
		Phone:         phone,
		Hospital:      hospital,
		Password:      password,
		Role:          "pharmacist",
		Verified:      false,
		Patients:      []string{""},
		List:          []string{""},
		Requests:      []string{""},
	}

	userAsBytes, _ := json.Marshal(pharmacist)

	return ctx.GetStub().PutState(userid, userAsBytes)
}

func (s *SmartContract) CreateTechnician(ctx contractapi.TransactionContextInterface, userid string, name string, email string, cnic string, phone string, hospital string, password string, id string) error {

	technician := LabTechnician{
		Technician_id: id,
		Name:          name,
		Email:         email,
		CNIC:          cnic,
		Phone:         phone,
		Hospital:      hospital,
		Password:      password,
		Role:          "labtechnician",
		Verified:      false,
		Patients:      []string{""},
		List:          []string{""},
		Requests:      []string{""},
	}

	userAsBytes, _ := json.Marshal(technician)

	return ctx.GetStub().PutState(userid, userAsBytes)
}

func (s *SmartContract) CreatePatient(ctx contractapi.TransactionContextInterface, userid string, name string, email string, cnic string, phone string, address string, city string, province string, age string, gender string, weight string, height string, password string, id string, photo string) error {

	patient := Patient{
		Name:     name,
		Email:    email,
		CNIC:     cnic,
		Phone:    phone,
		Photo:    photo,
		Password: password,
		Role:     "patient",
		Address:  address,
		City:     city,
		Province: province,
		Age:      age,
		Gender:   gender,
		Weight:   weight,
		Height:   height,
		Vitals:   "",
		Drugs:    []Drug{},
		List:     []string{""},
		Requests: []string{""},
		Reports:  []string{""},
	}

	patientAsBytes, _ := json.Marshal(patient)

	return ctx.GetStub().PutState(userid, patientAsBytes)
}

func (s *SmartContract) CreateAdmin(ctx contractapi.TransactionContextInterface, userid string, name string, email string, password string) error {

	admin := Admin{
		Email:    email,
		Password: password,
		Name:     name,
		Role:     "admin",
	}
	adminAsBytes, _ := json.Marshal(admin)

	return ctx.GetStub().PutState(userid, adminAsBytes)
}

func (s *SmartContract) CreateReport(ctx contractapi.TransactionContextInterface, reportid string, pID string, dname string, pname string, date string, symp string, advice string, vitals string) error {

	report := Report{
		Report_id: reportid,
		Doctor:    dname,
		Patient:   pname,
		Date:      date,
		Symptoms:  symp,
		Advice:    advice,
		Tests:     []Test{},
		Vitals:    vitals,
	}
	reportAsBytes, _ := json.Marshal(report)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.Reports = append(patient.Reports, reportid)

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)

	return ctx.GetStub().PutState(reportid, reportAsBytes)
}

func (s *SmartContract) CreateVitals(ctx contractapi.TransactionContextInterface, vitalsid string, pID string, date string, bloodpres string, temp string, weight string, height string) error {

	vitals := Vitals{
		Vitals_id: vitalsid,
		Patient:   pID,
		Date:      date,
		BloodPres: bloodpres,
		Temp:      temp,
		Weight:    weight,
		Height:    height,
	}
	vitalsAsBytes, _ := json.Marshal(vitals)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.Vitals = vitalsid

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)

	return ctx.GetStub().PutState(vitalsid, vitalsAsBytes)
}

func (s *SmartContract) CreateAppointment(ctx contractapi.TransactionContextInterface, appointmentid string, dname string, pname string, day string, time string) error {

	appointment := Appointment{
		Appointment_id: "Appointment",
		Doctor:         dname,
		Patient:        pname,
		Day:            day,
		Time:           time,
		Patients:       []string{""},
		List:           []string{""},
		Requests:       []string{""},
	}
	appointmentAsBytes, _ := json.Marshal(appointment)

	return ctx.GetStub().PutState(appointmentid, appointmentAsBytes)
}

//QueryDoctor returns the doctor stored in the world state with given id
func (s *SmartContract) QueryDoctor(ctx contractapi.TransactionContextInterface, name string) (*Doctor, error) {
	doctorAsBytes, err := ctx.GetStub().GetState(name)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if doctorAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", name)
	}

	doctor := new(Doctor)
	_ = json.Unmarshal(doctorAsBytes, doctor)

	return doctor, nil
}
func (s *SmartContract) QueryAdmin(ctx contractapi.TransactionContextInterface, name string, passw string) (bool, error) {
	adminAsBytes, err := ctx.GetStub().GetState("ADMIN1")

	if err != nil {
		return false, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if adminAsBytes == nil {
		return false, fmt.Errorf("%s does not exist", name)
	}
	var result = false
	admin := new(Admin)
	_ = json.Unmarshal(adminAsBytes, admin)
	if admin.Email == name {
		if admin.Password == passw {
			result = true
		}
	}
	return result, nil
}
func (s *SmartContract) QueryNurse(ctx contractapi.TransactionContextInterface, Email string) (*Nurse, error) {
	nurseAsBytes, err := ctx.GetStub().GetState(Email)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if nurseAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Email)
	}

	nurse := new(Nurse)
	_ = json.Unmarshal(nurseAsBytes, nurse)

	return nurse, nil
}
func (s *SmartContract) QueryTechnician(ctx contractapi.TransactionContextInterface, Email string) (*LabTechnician, error) {
	technicianAsBytes, err := ctx.GetStub().GetState(Email)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if technicianAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Email)
	}

	technician := new(LabTechnician)
	_ = json.Unmarshal(technicianAsBytes, technician)

	return technician, nil
}
func (s *SmartContract) QueryPharmacist(ctx contractapi.TransactionContextInterface, Email string) (*Pharmacist, error) {
	pharmacistAsBytes, err := ctx.GetStub().GetState(Email)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if pharmacistAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Email)
	}

	pharmacist := new(Pharmacist)
	_ = json.Unmarshal(pharmacistAsBytes, pharmacist)

	return pharmacist, nil
}

func (s *SmartContract) QueryPatient(ctx contractapi.TransactionContextInterface, name string) (*Patient, error) {
	patientAsBytes, err := ctx.GetStub().GetState(name)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if patientAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", name)
	}

	patient := new(Patient)
	_ = json.Unmarshal(patientAsBytes, patient)

	return patient, nil
}

func (s *SmartContract) QueryReport(ctx contractapi.TransactionContextInterface, name string) (*Report, error) {
	reportAsBytes, err := ctx.GetStub().GetState(name)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if reportAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", name)
	}

	report := new(Report)
	_ = json.Unmarshal(reportAsBytes, report)

	return report, nil
}

func (s *SmartContract) QueryVitals(ctx contractapi.TransactionContextInterface, name string) (*Vitals, error) {
	vitalsAsBytes, err := ctx.GetStub().GetState(name)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if vitalsAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", name)
	}

	vitals := new(Vitals)
	_ = json.Unmarshal(vitalsAsBytes, vitals)

	return vitals, nil
}

func (s *SmartContract) QueryAppointment(ctx contractapi.TransactionContextInterface, name string) (*Appointment, error) {
	appointmentAsBytes, err := ctx.GetStub().GetState(name)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if appointmentAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", name)
	}

	appointment := new(Appointment)
	_ = json.Unmarshal(appointmentAsBytes, appointment)

	return appointment, nil
}

// QueryUser returns the user stored in the world state with given id
func (s *SmartContract) QueryUser(ctx contractapi.TransactionContextInterface, Email string) (*User, error) {
	propertyAsBytes, err := ctx.GetStub().GetState(Email)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if propertyAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", Email)
	}

	user := new(User)
	_ = json.Unmarshal(propertyAsBytes, user)

	return user, nil
}

/////////////// Query All User //////////
func (s *SmartContract) QueryAllUsers(ctx contractapi.TransactionContextInterface) ([]QueryDoctor, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	fmt.Println(resultsIterator)
	fmt.Println(startKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryDoctor{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		doctor := new(Doctor)
		_ = json.Unmarshal([]byte(queryResponse.Value), doctor)

		queryDoctor := QueryDoctor{Key: queryResponse.Key, Record: doctor}
		results = append(results, queryDoctor)
	}

	return results, nil
}

///////////////Query All Doctors//////////
func (s *SmartContract) QueryAllDoctors(ctx contractapi.TransactionContextInterface) ([]QueryDoctor, error) {
	startKey := "DOCTOR0"
	endKey := "DOCTOR9999"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	fmt.Println(resultsIterator)
	fmt.Println(startKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryDoctor{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		doctor := new(Doctor)
		_ = json.Unmarshal([]byte(queryResponse.Value), doctor)

		queryDoctor := QueryDoctor{Key: queryResponse.Key, Record: doctor}
		results = append(results, queryDoctor)
	}

	return results, nil
}

///////////////Query All Nurses//////////
func (s *SmartContract) QueryAllNurses(ctx contractapi.TransactionContextInterface) ([]QueryNurse, error) {
	startKey := "NURSE0"
	endKey := "NURSE9999"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryNurse{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		nurse := new(Nurse)
		_ = json.Unmarshal([]byte(queryResponse.Value), nurse)

		queryNurse := QueryNurse{Key: queryResponse.Key, Record: nurse}
		results = append(results, queryNurse)
	}

	return results, nil
}

///////////////Query All Lab Technicians//////////
func (s *SmartContract) QueryAllTechnicians(ctx contractapi.TransactionContextInterface) ([]QueryTechnician, error) {
	startKey := "LTECHNICIAN0"
	endKey := "LTECHNICIAN9999"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryTechnician{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		technician := new(LabTechnician)
		_ = json.Unmarshal([]byte(queryResponse.Value), technician)

		queryTechnician := QueryTechnician{Key: queryResponse.Key, Record: technician}
		results = append(results, queryTechnician)
	}

	return results, nil
}

///////////////Query All Pharmacists//////////
func (s *SmartContract) QueryAllPharmacists(ctx contractapi.TransactionContextInterface) ([]QueryPharmacist, error) {
	startKey := "PHARMACIST0"
	endKey := "PHARMACIST9999"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryPharmacist{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		pharmacist := new(Pharmacist)
		_ = json.Unmarshal([]byte(queryResponse.Value), pharmacist)

		queryPharmacist := QueryPharmacist{Key: queryResponse.Key, Record: pharmacist}
		results = append(results, queryPharmacist)
	}

	return results, nil
}

///////////////Query All Patients//////////
func (s *SmartContract) QueryAllPatients(ctx contractapi.TransactionContextInterface) ([]QueryPatient, error) {
	startKey := "PATIENT0"
	endKey := "PATIENT9999"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryPatient{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		patient := new(Patient)
		_ = json.Unmarshal([]byte(queryResponse.Value), patient)

		queryPatient := QueryPatient{Key: queryResponse.Key, Record: patient}
		results = append(results, queryPatient)
	}

	return results, nil
}

///////////////Query All Appointments//////////
func (s *SmartContract) QueryAllAppointments(ctx contractapi.TransactionContextInterface) ([]QueryAppointment, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryAppointment{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		appointment := new(Appointment)
		_ = json.Unmarshal([]byte(queryResponse.Value), appointment)

		queryAppointment := QueryAppointment{Key: queryResponse.Key, Record: appointment}
		results = append(results, queryAppointment)
	}

	return results, nil
}

// Verify updates the verified field of staff with given id in world state
func (s *SmartContract) VerifyDoctor(ctx contractapi.TransactionContextInterface, name string) error {

	user, err := s.QueryDoctor(ctx, name)

	if err != nil {
		return err
	}

	user.Verified = true

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(name, userAsBytes)
}

func (s *SmartContract) VerifyNurse(ctx contractapi.TransactionContextInterface, name string) error {

	user, err := s.QueryNurse(ctx, name)

	if err != nil {
		return err
	}

	user.Verified = true

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(name, userAsBytes)
}

func (s *SmartContract) VerifyPharmacist(ctx contractapi.TransactionContextInterface, name string) error {

	user, err := s.QueryDoctor(ctx, name)

	if err != nil {
		return err
	}

	user.Verified = true

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(name, userAsBytes)
}

func (s *SmartContract) VerifyTechnician(ctx contractapi.TransactionContextInterface, name string) error {

	user, err := s.QueryTechnician(ctx, name)

	if err != nil {
		return err
	}

	user.Verified = true

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(name, userAsBytes)
}

//////////// Send Request ///////////
func (s *SmartContract) NewRequest(ctx contractapi.TransactionContextInterface, dID string, pID string, hosp string, pname string) error {

	doctor, err := s.QueryDoctor(ctx, dID)

	if err != nil {
		return err
	}

	doctor.Requests = append(doctor.Requests, pname)

	doctorAsBytes, _ := json.Marshal(doctor)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.Requests = append(patient.Requests, hosp)

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)
	return ctx.GetStub().PutState(dID, doctorAsBytes)
}

////////    Consent List /////////////
func (s *SmartContract) AddDocToConsent(ctx contractapi.TransactionContextInterface, dID string, pID string, dname string, pname string) error {

	doctor, err := s.QueryDoctor(ctx, dID)

	if err != nil {
		return err
	}

	doctor.Patients = append(doctor.Patients, pname)
	i := 1
	copy(doctor.Requests[i:], doctor.Requests[i+1:])
	doctor.Requests[len(doctor.Requests)-1] = ""
	doctor.Requests = doctor.Requests[:len(doctor.Requests)-1]

	doctorAsBytes, _ := json.Marshal(doctor)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.List = append(patient.List, dname)
	copy(patient.Requests[i:], patient.Requests[i+1:])
	patient.Requests[len(patient.Requests)-1] = ""
	patient.Requests = patient.Requests[:len(patient.Requests)-1]

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)
	return ctx.GetStub().PutState(dID, doctorAsBytes)
}

func (s *SmartContract) AddNurseToConsent(ctx contractapi.TransactionContextInterface, dID string, pID string, dname string, pname string) error {

	nurse, err := s.QueryNurse(ctx, dID)

	if err != nil {
		return err
	}

	nurse.Patients = append(nurse.Patients, pname)

	nurseAsBytes, _ := json.Marshal(nurse)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.List = append(patient.List, dname)

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)
	return ctx.GetStub().PutState(dID, nurseAsBytes)
}

func (s *SmartContract) AddPharToConsent(ctx contractapi.TransactionContextInterface, dID string, pID string, dname string, pname string) error {

	pharmacist, err := s.QueryPharmacist(ctx, dID)

	if err != nil {
		return err
	}

	pharmacist.Patients = append(pharmacist.Patients, pname)

	pharmacistAsBytes, _ := json.Marshal(pharmacist)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.List = append(patient.List, dname)

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)
	return ctx.GetStub().PutState(dID, pharmacistAsBytes)
}

func (s *SmartContract) AddLabToConsent(ctx contractapi.TransactionContextInterface, dID string, pID string, dname string, pname string) error {

	lab, err := s.QueryTechnician(ctx, dID)

	if err != nil {
		return err
	}

	lab.Patients = append(lab.Patients, pname)

	labAsBytes, _ := json.Marshal(lab)

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.List = append(patient.List, dname)

	patientAsBytes, _ := json.Marshal(patient)

	ctx.GetStub().PutState(pID, patientAsBytes)
	return ctx.GetStub().PutState(dID, labAsBytes)
}

///////////// Prescribe Drug /////////
func (s *SmartContract) NewDrugRequest(ctx contractapi.TransactionContextInterface, pID string, dname string) error {

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	patient.Requests = append(patient.Requests, "*"+dname)

	patientAsBytes, _ := json.Marshal(patient)

	return ctx.GetStub().PutState(pID, patientAsBytes)
}

func (s *SmartContract) AprroveDrugRequest(ctx contractapi.TransactionContextInterface, pID string, dname string) error {

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	for i := 0; i < len(patient.Drugs); i++ {
		if dname == (patient.Drugs[i].Name) {
			patient.Drugs[i].Given = true
		}
	}
	patientAsBytes, _ := json.Marshal(patient)

	return ctx.GetStub().PutState(pID, patientAsBytes)
}

func (s *SmartContract) AddDrug(ctx contractapi.TransactionContextInterface, pID string, name string, time string, meal string, duration string, date string) error {

	patient, err := s.QueryPatient(ctx, pID)

	if err != nil {
		return err
	}

	drug := Drug{
		Name:     name,
		Timing:   time,
		Duration: duration,
		Meal:     meal,
		Date:     date,
		Given:    false,
	}

	patient.Drugs = append(patient.Drugs, drug)

	patientAsBytes, _ := json.Marshal(patient)

	// patient, err := s.QueryPatient(ctx, pID)

	// if err != nil {
	// 	return err
	// }

	// patient.List = append(patient.List, dname)

	// patientAsBytes, _ := json.Marshal(patient)

	// ctx.GetStub().PutState(pID, patientAsBytes)
	return ctx.GetStub().PutState(pID, patientAsBytes)
}

func (s *SmartContract) AddTest(ctx contractapi.TransactionContextInterface, rID string, name string) error {

	report, err := s.QueryReport(ctx, rID)

	if err != nil {
		return err
	}

	test := Test{
		Name: name,
		Path: "",
	}

	report.Tests = append(report.Tests, test)

	reportAsBytes, _ := json.Marshal(report)

	return ctx.GetStub().PutState(rID, reportAsBytes)
}

func (s *SmartContract) UploadTest(ctx contractapi.TransactionContextInterface, rID string, tname string, path string) error {

	report, err := s.QueryReport(ctx, rID)

	if err != nil {
		return err
	}

	for i := 0; i < len(report.Tests); i++ {
		if tname == (report.Tests[i].Name) {
			report.Tests[i].Path = path
		}
	}
	reportAsBytes, _ := json.Marshal(report)

	return ctx.GetStub().PutState(rID, reportAsBytes)
}

func (s *SmartContract) CheckUser(ctx contractapi.TransactionContextInterface, Email string, password string) bool {
	user, err := s.QueryUser(ctx, Email)

	if err != nil {
		return true
	}
	if user.Password == password {
		fmt.Printf("Successfull Login: %s", Email)
		return true
	}

	return false
}

//Checks if the registered user is an admin or not
func (s *SmartContract) CheckUser1(ctx contractapi.TransactionContextInterface, Email string) bool {
	user, err := s.QueryUser(ctx, Email)

	if err != nil {
		return true
	}
	if user.IsAdmin == false {
		fmt.Printf("Successfull Login: %s", Email)
		return false
	}

	return true
}
func (s *SmartContract) GetUserId(ctx contractapi.TransactionContextInterface, Email string) string {
	user, err := s.QueryUser(ctx, Email)

	if err != nil {
		return "false"
	}
	var x = user.CNIC
	return x
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabproperty chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabproperty chaincode: %s", err.Error())
	}
}
