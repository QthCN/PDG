function Config() {
    if (process.env.NODE_ENV === "development") {
        this.baseAddress = "http://127.0.0.1:18080"
    } else {
        this.baseAddress = ""
    }
    
}

Config.prototype.getAddress = function(target) {
    var url = ""
    switch (target) {
        case "LOGIN_PAGE":
                url = `${this.baseAddress}/login.html`
                break
        case "GET_USERNAME":
            url = `${this.baseAddress}/v1/ajax/auth/info`
            break
        default:
            console.error(`unknown target: ${target}`)
            break
    }
    return url
}

export default Config