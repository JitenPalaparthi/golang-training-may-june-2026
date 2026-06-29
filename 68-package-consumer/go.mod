module demo

go 1.26.3

require github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes v0.0.0-20260629142136-16a39ce69300 // indirect

replace github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes v0.0.0-20260629142136-16a39ce69300 => ../golang-personal-training-june-july-2026-shapes

// replace is only used for the development and testing .. later it has to be properly versioned and pushed
// the policy is based on the org or org