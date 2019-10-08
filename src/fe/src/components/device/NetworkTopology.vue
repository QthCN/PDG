<template>
  <div style="width: 100%; height: 100%">
      <div id="mountNode"></div>
  </div>
</template>

<script>
import axios from "axios"
import G6 from '@antv/g6'

import Config from '../../config'

import logoNetworkRouter0 from '../../assets/logo-network-router-0.png'
import logoNetworkRouter1 from '../../assets/logo-network-router-1.png'
import logoServer0 from '../../assets/logo-server-0.png'



export default {
  name: 'NetworkTopology',
  data () {
      return {
          config: new Config(),
          connections: [],
      }
  },
  computed: {
      networkData: function() {
          var that = this
          
          // 帮助函数
          var deviceExist = function(uuid, devices) {
              for (var device of devices) {
                  if (device.uuid === uuid) {
                      return true
                  }
              }
              return false
          }

          var getConnectedDevices = function(uuid) {
              var devices = []
              for (var connection of that.connections) {
                  if (connection.source_id === uuid) {
                      devices.push( {
                            uuid: connection.destination_id,
                            name: connection.destination_device_name,
                            deviceType: connection.destination_device_type,
                      })
                  } else if (connection.destination_id === uuid) {
                      devices.push( {
                            uuid: connection.source_id,
                            name: connection.source_device_name,
                            deviceType: connection.source_device_type,
                      })
                  } else {
                      
                  }
              }
              return devices
          }

          var items = {nodes: [], edges: []}

          // 根据连接信息计算拓扑，非Network的在最下层，然后根据网络设备往上推导计算
          var layers = []
          var devicesAnalysised = []

          // 首先获取所有的非网络设备
          var nonNetworkDevices = []
          for (var connection of that.connections) {
              // 只考虑连接网络设备的情况
              if (connection.source_device_type !== "NETWORK" && connection.destination_device_type !== "NETWORK") {
                  continue
              }

              if (connection.source_device_type !== "NETWORK" && deviceExist(connection.source_id, nonNetworkDevices) === false) {
                  nonNetworkDevices.push({
                      uuid: connection.source_id,
                      name: connection.source_device_name,
                      deviceType: connection.source_device_type,
                  })
                  devicesAnalysised.push({
                      uuid: connection.source_id,
                      name: connection.source_device_name,
                      deviceType: connection.source_device_type,
                  })
              }

              if (connection.destination_device_type !== "NETWORK" && deviceExist(connection.destination_id, nonNetworkDevices) === false) {
                  nonNetworkDevices.push({
                      uuid: connection.destination_id,
                      name: connection.destination_device_name,
                      deviceType: connection.destination_device_type,
                  })
                  devicesAnalysised.push({
                      uuid: connection.destination_id,
                      name: connection.destination_device_name,
                      deviceType: connection.destination_device_type,
                  })
              }
          }
          // 将这些非网络设备推入layer
          layers.push(nonNetworkDevices)

          // 根据layers的最后一层计算连接的网络设备，如果为空则停止
          var calcDevices = nonNetworkDevices
          while (calcDevices.length !== 0) {
              calcDevices = []
              var toCalcDevices = layers[layers.length - 1]
              for (var toCalcDevice of toCalcDevices) {
                  var connectedDevices = getConnectedDevices(toCalcDevice.uuid)
                  for (var connectedDevice of connectedDevices) {
                      if (connectedDevice.deviceType != "NETWORK") {
                          continue
                      }
                      if (deviceExist(connectedDevice.uuid, devicesAnalysised) === true) {
                          continue
                      }

                      calcDevices.push(connectedDevice)
                      devicesAnalysised.push(connectedDevice)
                  }
              }
              if (calcDevices.length !== 0) {
                  layers.push(calcDevices)
              }
          }

          // 计算结果数据
          // nodes
          var xOffset = 200
          var layerNum = layers.length
          var layerMaxItems = 0
          for (var layer of layers) {
              if (layer.length > layerMaxItems) {
                  layerMaxItems = layer.length 
              }
          }
          var layerHeight = 100 * layerMaxItems

          layers.reverse()
          var idx = 1
          for (var layer of layers) {
              var yOffset = layerHeight / (layer.length + 1)
              var itemIdx = 1
              for (var item of layer) {
                  var image = logoNetworkRouter0
                  if (item.deviceType !== "NETWORK") {
                      image = logoServer0
                  }
                  if (idx === 1 && item.deviceType === "NETWORK") {
                      image = logoNetworkRouter1
                  }

                  var shape = 'image'
                  // 制造告警效果
                  if (idx === itemIdx) {
                      shape = 'image-red'
                  }
                  items.nodes.push({
                      id: item.uuid,
                      shape: shape,
                      label: item.name,
                      img: image,
                      x: idx * xOffset,
                      y: itemIdx * yOffset,
                  })

                  itemIdx += 1
              }
              idx += 1
          }

          // edges
          for (var layer of layers) {
              for (var item of layer) {
                  var connectedDevices = getConnectedDevices(item.uuid)
                  for (var connectedDevice of connectedDevices) {
                      // 判断是否已经存在此记录
                      var recordExist = false
                      for (var edge of items.edges) {
                          if ((edge.source === item.uuid && edge.target === connectedDevice.uuid) || (edge.target === item.uuid && edge.source === connectedDevice.uuid)) {
                              recordExist = true
                              break
                          }
                      }
                      if (recordExist === false) {
                          items.edges.push({
                              shape: "line-running",
                              source: item.uuid,
                              target: connectedDevice.uuid,
                          })
                      }
                  }
              }
          }

          console.log(JSON.stringify(items))
          return items
      }
  },
  created () {
    var that = this
    that.initG6()
    that.initData()
  },
  components: {
  },
  mounted () {
    
  },
  methods: {
    initG6() {
        G6.registerEdge('line-running', {
            afterDraw(cfg, group) {
                const shape = group.get('children')[0];
                const startPoint = shape.getPoint(0);
                // 添加圆点
                const circle = group.addShape('circle', {
                    attrs: {
                        x: startPoint.x,
                        y: startPoint.y,
                        fill: 'blue',
                        r: 3
                    }
                });
                
                // 对圆点添加动画
                circle.animate({
                    onFrame(ratio) {
                        const tmpPoint = shape.getPoint(1.0-ratio);
                        return {
                            x: tmpPoint.x,
                            y: tmpPoint.y
                        };
                    },
                    repeat: true
                }, 3000);
            }
        }, 'line');

        G6.registerNode('image-red', {
            afterDraw(cfg, group) {
                const r = 25;
                const back1 = group.addShape('circle',{
                    zIndex: -3,
                    attrs: {
                        x: 0,
                        y: 0,
                        r,
                        fill: 'red',
                        opacity: 0.6
                    }
                });
                const back2 = group.addShape('circle',{
                    zIndex: -2,
                    attrs: {
                        x: 0,
                        y: 0,
                        r,
                        fill: 'red', // 为了显示清晰，随意设置了颜色
                        opacity: 0.6
                    }
                });

                const back3 = group.addShape('circle',{
                    zIndex: -1,
                    attrs: {
                        x: 0,
                        y: 0,
                        r,
                        fill: 'red',
                        opacity: 0.6
                    }
                });
                group.sort(); // 排序，根据zIndex 排序
                
                back1.animate({ // 逐渐放大，并消失
                    r: r + 10,
                    opacity: 0.1,
                    repeat: true // 循环
                }, 3000, 'easeCubic', null, 0) // 无延迟

                back2.animate({ // 逐渐放大，并消失
                    r: r + 10,
                    opacity: 0.1,
                    repeat: true // 循环
                }, 3000, 'easeCubic', null, 1000) // 1 秒延迟

                back3.animate({ // 逐渐放大，并消失
                    r: r + 10,
                    opacity: 0.1,
                    repeat: true // 循环
                }, 3000, 'easeCubic', null, 2000) // 2 秒延迟
            }
        }, 'image');
    },
    initData () {
        var that = this

        that.connections = []
        
        Promise.all([
            that.syncConnections()
        ]).then(values => {
            that.render()
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message.error("页面加载异常")
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    getActualWidth()
    {
        var actualWidth = window.innerWidth ||
                        document.documentElement.clientWidth ||
                        document.body.clientWidth ||
                        document.body.offsetWidth;

        return actualWidth;
    },
    getActualHeight()
    {
        var actualHeight = window.innerHeight ||
                        document.documentElement.clientHeight ||
                        document.body.clientHeight ||
                        document.body.offsetHeight;
        return actualHeight;
    },
    render () {
        var that = this
        const graph = new G6.Graph({
            container: 'mountNode',
            fitView: true,
            fitViewPadding: [ 20, 20, 20, 20 ],
            width: that.getActualWidth() - 250,
            height: that.getActualHeight() - 70,
            nodeStyle: {
                default: {
                    fill: '#40a9ff',
                    stroke: '#096dd9',
                    labelCfg: {
                        position: 'bottom'
                    }
                }
            },
            edgeStyle: {
                default: { 
                    stroke: '#A3B1BF' ,
                    lineWidth: 3,
                }
            },
            modes: {
                default: [ 'drag-canvas', 'zoom-canvas' ]  // 允许拖拽画布、放缩画布
            }
            });
        graph.read(that.networkData);

        graph.on('node:click', e => {
            var deviceInfo = e.item._cfg.model.deviceInfo
        })

        graph.on('edge:click', e => {
            var connectionInfo = e.item._cfg.model.connectionInfo
        })
    },
    syncConnections () {
        var that = this
        return axios.post(that.config.getAddress("LIST_CONNECTIONS"))
                    .then(response => {
                        that.connections = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.connections = []
                        that.$message.error("获取数据异常")
                    })
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
