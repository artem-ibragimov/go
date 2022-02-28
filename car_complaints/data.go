package car_complaints

// type Brand struct {
// 	name   string
// 	url    string
// 	models []Model
// }

// type Model struct {
// 	name string
// 	url  string
// 	ages []ProblemAge
// }

// type ProblemAge struct {
// 	// год жалоб
// 	year uint
// 	// кол-во жалоб за этот год
// 	count      uint
// 	categories []ProblemCategory
// }

// type ProblemCategory struct {
// 	name     string
// 	count    uint
// 	problems []Problem
// }

type Problem struct {
	brand           string
	model           string
	year            uint
	category        string
	name            string
	count           uint
	severity_rating uint8
	repair_cost     uint
	average_mileage uint
}
