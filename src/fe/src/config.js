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
        case "LIST_USERS":
                url = `${this.baseAddress}/v1/ajax/auth/user/list`
                break
        case "REMOVE_USER":
                url = `${this.baseAddress}/v1/ajax/auth/user/remove`
                break     
        case "CREATE_USER":
                url = `${this.baseAddress}/v1/ajax/auth/user/create`
                break     
        default:
            console.error(`unknown target: ${target}`)
            break
    }
    return url
}

export default Config