import axios from 'axios'
import Config from '../config'

function Auth() {
    this.config = new Config()
}

Auth.prototype.getCurrentUsername = function(callback) {
    var that = this

    if (process.env.NODE_ENV === "development") {
        callback("ADMIN")
    }

    axios.post(that.config.getAddress("GET_USERNAME"))
        .then(function(response){
            if (response.data.username !== "") {
                callback(response.data.username)
            } else {
                window.location.href = that.config.getAddress("LOGIN_PAGE")
            }
        })
        .catch(function(error){
            console.error(JSON.stringify(error))
            window.location.href = that.config.getAddress("LOGIN_PAGE")
        })
}

export default Auth