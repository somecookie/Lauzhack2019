package blockchain

//Block represents one block of the block chain.
//FirstName  string is the first name of the doctor in charge.
//Name       string is the name of the doctor.
//Location   string is the name of the hospital.
//isSuccess    bool tells if the operation was a success or not.
//ReportHash string is the hash of the report written after the operation.
//Total      int is the total number of operations made by the doctor
//Succeed    int is the number of operations that were a success
//SuccessRate float32 is defined as Succeed/Total
type Block struct {
	FirstName   string
	Name        string
	Location    string
	Success     bool
	ReportHash  string
	Total       int
	Succeed     int
	SuccessRate float64
}
