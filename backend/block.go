package backend

//Block represents one block of the block chain.
//FirstName  string is the first name of the doctor in charge.
//Name       string is the name of the doctor.
//Success    bool tells if the operation was a success or not.
//ReportHash string is the hash of the report written after the operation.
type Block struct {
	FirstName  string
	Name       string
	Success    bool
	ReportHash string
}

