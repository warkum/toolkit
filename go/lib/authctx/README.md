# Auth Context

get information authentication user from context

# How to use

Here some sample code to get user identifier
```
func ActionNeedAuthenticate(ctx context.Context) (err error) {
	userID := authctx.GetUserIDInContext(ctx)
	if userID == "" {
		return errors.New("fail exec this function due to missing user authorization")
	}
    log.Println("User ID: ", userID)
    return nil
}
```

here sample code to get user level
```
func ActionNeedAuthenticate(ctx context.Context) (err error) {
	levelUser := authctx.GetUserLevelInContext(ctx)
	if levelUser == "" {
		return errors.New("fail exec this function due to missing user authorization")
	}

    log.Println("Level User: ", levelUser)
    return nil
}
```

We have 3 level of users.
```
	AUTH_ADMIN AUTH_LEVEL = "ADMINISTRATOR"
	AUTH_OPS   AUTH_LEVEL = "OPS"
	AUTH_BUYER AUTH_LEVEL = "BUYER"
```


