func APICall(ctx context.Context, args) {

  // check for cancellation
  if ctx.Done() {
    return 
  }

  // call API in goroutine 
  go func() {
    result := thirdPartyAPI(args) 
    
    // check for cancellation
    if ctx.Done() {
      return
    }

    handleResults(result)
  }()
}
