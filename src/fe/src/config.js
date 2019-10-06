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
        case "GET_PHYSICAL_TOPOLOGY":
                url = `${this.baseAddress}/v1/ajax/device/topology/physical`
                break
        case "GET_RESOURCE_TOPOLOGY":
                url = `${this.baseAddress}/v1/ajax/device/topology/resource`
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
        case "LIST_DATACENTERS":
                url = `${this.baseAddress}/v1/ajax/device/datacenter/list`
                break
        case "DELETE_DATACENTER":
                url = `${this.baseAddress}/v1/ajax/device/datacenter/remove`
                break     
        case "CREATE_DATACENTER":
                url = `${this.baseAddress}/v1/ajax/device/datacenter/create`
                break 
        case "LIST_RACKS":
                url = `${this.baseAddress}/v1/ajax/device/rack/list`
                break
        case "MAPPING_DEVICE_RACK":
                url = `${this.baseAddress}/v1/ajax/device/map/rack`
                break
        case "MAPPING_RACK_DATACENTER":
                url = `${this.baseAddress}/v1/ajax/device/rack/map/datacenter`
                break
        case "DELETE_RACK":
                url = `${this.baseAddress}/v1/ajax/device/rack/remove`
                break     
        case "CREATE_RACK":
                url = `${this.baseAddress}/v1/ajax/device/rack/create`
                break 
        case "LIST_SERVERS":
                url = `${this.baseAddress}/v1/ajax/device/server/list`
                break
        case "DELETE_SERVER":
                url = `${this.baseAddress}/v1/ajax/device/server/remove`
                break     
        case "CREATE_SERVER":
                url = `${this.baseAddress}/v1/ajax/device/server/create`
                break 
        case "LIST_NETWORK_DEVICES":
                url = `${this.baseAddress}/v1/ajax/device/network/list`
                break
        case "DELETE_NETWORK_DEVICE":
                url = `${this.baseAddress}/v1/ajax/device/network/remove`
                break     
        case "CREATE_NETWORK_DEVICE":
                url = `${this.baseAddress}/v1/ajax/device/network/create`
                break 
        case "LIST_STORAGE_DEVICES":
                url = `${this.baseAddress}/v1/ajax/device/storage/list`
                break
        case "DELETE_STORAGE_DEVICE":
                url = `${this.baseAddress}/v1/ajax/device/storage/remove`
                break     
        case "CREATE_STORAGE_DEVICE":
                url = `${this.baseAddress}/v1/ajax/device/storage/create`
                break 
        case "LIST_COMMON_DEVICES":
                url = `${this.baseAddress}/v1/ajax/device/common/list`
                break
        case "DELETE_COMMON_DEVICE":
                url = `${this.baseAddress}/v1/ajax/device/common/remove`
                break     
        case "CREATE_COMMON_DEVICE":
                url = `${this.baseAddress}/v1/ajax/device/common/create`
                break 
        case "LIST_IPS":
                url = `${this.baseAddress}/v1/ajax/ip/list`
                break 
        case "CREATE_IP":
                url = `${this.baseAddress}/v1/ajax/ip/create`
                break 
        case "DELETE_IP":
                url = `${this.baseAddress}/v1/ajax/ip/remove`
                break 
        case "LIST_IPSETS":
                url = `${this.baseAddress}/v1/ajax/ipset/list`
                break 
        case "CREATE_IPSET":
                url = `${this.baseAddress}/v1/ajax/ipset/create`
                break 
        case "DELETE_IPSET":
                url = `${this.baseAddress}/v1/ajax/ipset/remove`
                break 
        case "LIST_CONNECTIONS":
                url = `${this.baseAddress}/v1/ajax/connection/list`
                break 
        case "CREATE_CONNECTION":
                url = `${this.baseAddress}/v1/ajax/connection/create`
                break 
        case "DELETE_CONNECTION":
                url = `${this.baseAddress}/v1/ajax/connection/remove`
                break 
        default:
            console.error(`unknown target: ${target}`)
            break
    }
    return url
}

export default Config