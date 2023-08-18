package jwt

//var jwtKey = []byte("my_secret_key")
//
//type User struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//
//type Claims struct {
//	Username string `json:"username"`
//	jwt.StandardClaims
//}
//
//func main() {
//	router := mux.NewRouter()
//
//	router.HandleFunc("/login", login).Methods("POST")
//	router.HandleFunc("/verify", verify).Methods("GET")
//
//	log.Fatal(http.ListenAndServe(":8000", router))
//}
//
//func login(w http.ResponseWriter, r *http.Request) {
//	var user User
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// TODO: 在此处添加用户认证逻辑
//	expirationTime := time.Now().Add(5 * time.Minute)
//	claims := &Claims{
//		Username: user.Username,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtKey)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	http.SetCookie(w, &http.Cookie{
//		Name:    "token",
//		Value:   tokenString,
//		Expires: expirationTime,
//	})
//
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprintf(w, "登录成功")
//}
//
//func verify(w http.ResponseWriter, r *http.Request) {
//	c, err := r.Cookie("token")
//	if err != nil {
//		if err == http.ErrNoCookie {
//			http.Error(w, "未登录", http.StatusUnauthorized)
//			return
//		}
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	tokenString := c.Value
//
//	claims := &Claims{}
//
//	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//		return jwtKey, nil
//	})
//
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			http.Error(w, "无效的签名", http.StatusUnauthorized)
//			return
//		}
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if !token.Valid {
//		http.Error(w, "无效的令牌", http.StatusUnauthorized)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprintf(w, "令牌有效")
//}
