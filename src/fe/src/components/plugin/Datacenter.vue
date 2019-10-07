<template>
  <div style="">
    <canvas id="renderCanvas" touch-action="none" style="width: 100%; height: 100%;"></canvas> 

    <el-dialog title="设备信息" :visible.sync="serverDeviceStatusDialogVisible" width="80%">
        <DeviceStatus :uuid="deviceUUID" :device-type="deviceType"></DeviceStatus>
    </el-dialog>

    <el-dialog title="线路信息" :visible.sync="connectionStatusDialogVisible" width="80%">

    </el-dialog>
  </div>
</template>

<script>
import axios from "axios"
import * as BABYLON from 'babylonjs'

import Config from '../../config'

import rackfacePic from '../../assets/rackface0.png'
import serverfacePicIBM0 from '../../assets/serverface-ibm-0.png'
import networkfacePicCisco0 from '../../assets/networkface-cisco-0.png'
import storagefacePicIBM0 from '../../assets/storageface-ibm-0.png'
import floor0 from '../../assets/floor0.png'

import DeviceStatus from './DeviceStatus.vue'

export default {
  name: 'Datacenter',
  props: ['datacenter'],
  data () {
      return {
        serverDeviceStatusDialogVisible: false,
        deviceUUID: "",
        deviceType: "",

        connectionStatusDialogVisible: false,
      }
  },
  created () {

  },
  components: {
      DeviceStatus
  },
  mounted () {
    var that = this
    that.render()
  },
  methods: {
    showServerInfo (evt) {
        var server = evt.source.__server
        this.deviceUUID = server.uuid
        this.deviceType = server.type
        this.serverDeviceStatusDialogVisible = true
    },
    showConnectionInfo (evt) {
        var connection = evt.source.__connection
        this.connectionStatusDialogVisible = true
    },
    render () {
        var that = this

        var canvas = document.getElementById("renderCanvas"); // Get the canvas element 
        var engine = new BABYLON.Engine(canvas, true); // Generate the BABYLON 3D engine

        /******* Add the create scene function ******/
        var createScene = function () {

            // Create the scene space
            var scene = new BABYLON.Scene(engine);

            // Add a camera to the scene and attach it to the canvas
            var camera = new BABYLON.ArcRotateCamera("Camera", Math.PI * 1 - 0.6, Math.PI * 0.3, 120, new BABYLON.Vector3.Zero(), scene);
            camera.attachControl(canvas, true);

             // This creates a light
            var light = new BABYLON.HemisphericLight("light", new BABYLON.Vector3(-1, 1, 0), scene);

            // Default intensity is 1. Let's dim the light a small amount
            light.intensity = 0.7;

            that.doRender(scene, light)

            return scene;
        };
        /******* End of the create scene function ******/    

        var scene = createScene(); //Call the createScene function

        // Register a render loop to repeatedly render the scene
        engine.runRenderLoop(function () { 
                scene.render();
        });

        // Watch for browser/canvas resize events
        window.addEventListener("resize", function () { 
                engine.resize();
        });
    },

    /*
        机柜的长和宽都为60厘米，高根据U数的不同而不同，1U约为5厘米。因此都以1U为单位，即长宽为12单位。
    */
    doRender (scene, light) {
        var that = this

        // 绘制机房地面
        var ground = BABYLON.MeshBuilder.CreateGround("ground", {width: that.datacenter.size.width, height: that.datacenter.size.height}, scene);
        var groundMaterial = new BABYLON.StandardMaterial("groundmat", scene);
        groundMaterial.diffuseTexture = new BABYLON.Texture(floor0, scene);
        groundMaterial.diffuseTexture.uScale = 10; //you could try changin this and below value to see replicated image 
        groundMaterial.diffuseTexture.vScale = 10;
        groundMaterial.diffuseTexture.level = 1; //It is kind of z-index
        ground.material = groundMaterial;
        
        // 绘制机柜
        for (var rack of that.datacenter.racks) {
            // 材质
            var rackMat = new BABYLON.StandardMaterial("rackMat", scene)
            var rackTexture = new BABYLON.Texture(rackfacePic, scene)
            rackMat.diffuseTexture = rackTexture

            var rackFaceUV = new Array(6)
            for (var i = 0; i < 6; i++) {
                rackFaceUV[i] = new BABYLON.Vector4(0, 0, 0, 0);
            }
            rackFaceUV[3] = new BABYLON.Vector4(0, 0, 1, 1)

            // box
            let width = 12
            let depth = 12
            let height = rack.u
            var rackBox = BABYLON.MeshBuilder.CreateBox(`rackBox-${rack.name}`, {height: height, width: width, depth: depth, faceUV: rackFaceUV}, scene);
            rackBox.position.x = rack.x
            rackBox.position.z = rack.z
            rackBox.position.y = height / 2
            rackBox.visibility = 0.2
            rackBox.material = rackMat
            rackBox.isPickable = false

            
            // 服务器信息
            for (var server of rack.servers) {
                // 材质
                var serverMat = new BABYLON.StandardMaterial("serverMat", scene)
                var serverFaceColors = new Array(6);
                
                switch (server.status) {
                    case "GOOD":
                        serverFaceColors[0]= new BABYLON.Color3.Green
                        serverFaceColors[1]= new BABYLON.Color3.Green
                        serverFaceColors[2]= new BABYLON.Color3.Green
                        serverFaceColors[4]= new BABYLON.Color3.Green
                        serverFaceColors[5]= new BABYLON.Color3.Green
                        serverFaceColors[6]= new BABYLON.Color3.Green
                        break;
                    case "BAD":
                        serverFaceColors[0]= new BABYLON.Color3.Red
                        serverFaceColors[1]= new BABYLON.Color3.Red
                        serverFaceColors[2]= new BABYLON.Color3.Red
                        serverFaceColors[4]= new BABYLON.Color3.Red
                        serverFaceColors[5]= new BABYLON.Color3.Red
                        serverFaceColors[6]= new BABYLON.Color3.Red
                        break;
                    default:
                        break;
                }
                var serverTexture
                switch (server.type) {
                    case "SREVER":
                        serverTexture = new BABYLON.Texture(serverfacePicIBM0, scene)
                        break;

                    case "NETWORK":
                        serverTexture = new BABYLON.Texture(networkfacePicCisco0, scene)
                        break;

                    case "STORAGE":
                        serverTexture = new BABYLON.Texture(storagefacePicIBM0, scene)
                        break;
                
                    default:
                        serverTexture = new BABYLON.Texture(serverfacePicIBM0, scene)
                        break;
                }
                serverMat.diffuseTexture = serverTexture

                var serverFaceUV = new Array(6)
                for (var i = 0; i < 6; i++) {
                    serverFaceUV[i] = new BABYLON.Vector4(0, 0, 0, 0);
                }
                serverFaceUV[3] = new BABYLON.Vector4(0, 0, 1, 1)

                let width = 12 - 0.1
                let depth = 12 - 0.1
                let height = server.sizeU - 0.1
                var serverBox = BABYLON.MeshBuilder.CreateBox(`serverBox-${server.name}`, {height: height, width: width, depth: depth, faceUV: serverFaceUV, faceColors:serverFaceColors}, scene);
                serverBox.position.x = rack.x - 0.1
                serverBox.position.z = rack.z
                serverBox.position.y = server.begU + server.sizeU/2
                serverBox.material = serverMat

                serverBox.__server = server

                // 绑定点击方法
                serverBox.actionManager = new BABYLON.ActionManager(scene)
                serverBox.actionManager.registerAction(
                    new BABYLON.ExecuteCodeAction(
                        BABYLON.ActionManager.OnPickTrigger, 
                        that.showServerInfo
                    )
                )

                // 网络设备到其它设备的连线
                for (var connection of server.connections) {
                    var points = []

                    var z_offset = (connection.begU + connection.sizeU/2 - rack.u/2) / rack.u * 12

                    // 源设备直连点
                    var pointA_X = rack.x + 5.95
                    var pointA_Y = serverBox.position.y
                    var pointA_Z = rack.z + z_offset

                    points.push(new BABYLON.Vector3(pointA_X, pointA_Y, pointA_Z))

                    // 源设备拐点
                    var pointB_X = pointA_X + 2
                    var pointB_Y = serverBox.position.y
                    var pointB_Z = rack.z + z_offset

                    points.push(new BABYLON.Vector3(pointB_X, pointB_Y, pointB_Z))

                    // 目的设备拐点
                    var pointC_X = pointB_X
                    var pointC_Y = connection.begU + connection.sizeU/2
                    var pointC_Z = rack.z + z_offset

                    points.push(new BABYLON.Vector3(pointC_X, pointC_Y, pointC_Z))

                    // 目的设备直连点
                    var pointD_X = pointC_X - 2
                    var pointD_Y = connection.begU + connection.sizeU/2
                    var pointD_Z = rack.z + z_offset

                    points.push(new BABYLON.Vector3(pointD_X, pointD_Y, pointD_Z))

                    var lines = BABYLON.MeshBuilder.CreateLines("lines", {points: points}, scene)

                    lines.__connection = {
                        source: server,
                        destionation: connection,
                    }

                    lines.actionManager = new BABYLON.ActionManager(scene)
                    lines.actionManager.registerAction(
                        new BABYLON.ExecuteCodeAction(
                            BABYLON.ActionManager.OnPickTrigger, 
                            that.showConnectionInfo
                        )
                    ) 
                }
            }
        }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
