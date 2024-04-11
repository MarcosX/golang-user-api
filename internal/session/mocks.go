package session

// claim with userId set to 0
var ValidSession = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwIn0.hR9oaCu7Ud_Mr-QENEc-K6DLdZBaReap1rpvgnyEPU0"

// claim with userId set to 0
var InvalidUserValidSession = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI5OSJ9.PTxPX1DdhGraFnQp163hoXpafW0V-a-YOu55eWxblpA"

// jwt signed with wrong secret
var WrongSecretSession = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwIn0.nL3zQibYzBCqwzILJ6KJQSiYEEXjxqnu5rM0_U-ZH0E"
