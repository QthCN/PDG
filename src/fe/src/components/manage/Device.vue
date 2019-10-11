<template>
  <div>
    <el-tabs type="border-card">
        <el-tab-pane label="机房">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createDataCenterDialogVisible = true">新建机房</el-button>
            <el-table
                :data="datacenters"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    prop="name"
                    label="机房名">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="100">
                    <template slot-scope="scope">
                        <el-button @click="removeDataCenter(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建机房" :visible.sync="createDataCenterDialogVisible">
                <el-form :model="createDataCenterForm">
                    <el-form-item label="机房名" :label-width="formLabelWidth">
                        <el-input v-model="createDataCenterForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createDataCenterDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createDataCenter">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="机柜">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createRackDialogVisible = true">新建机柜</el-button>
            <el-table
                :data="racks"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    prop="name"
                    label="机柜名">
                </el-table-column>
                <el-table-column
                    prop="datacenter"
                    label="机房信息">
                </el-table-column>
                <el-table-column
                    prop="size_u"
                    label="总U数">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="200">
                    <template slot-scope="scope">
                        <el-button @click="showEditRackDatacenterDialog(scope.row)" type="primary" plain size="small">编辑机房信息</el-button>
                        <el-button @click="removeRack(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-dialog title="编辑机房信息" :visible.sync="editRackDatacenterDialogVisible">
                <el-form :model="editRackDatacenterForm">
                    <el-form-item label="机房" :label-width="formLabelWidth">
                        <el-select v-model="editRackDatacenterForm.datacenterId" placeholder="请选择">
                            <el-option
                                v-for="item in datacenters"
                                :key="item.uuid"
                                :label="item.name"
                                :value="item.uuid">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="X坐标" :label-width="formLabelWidth">
                        <el-input v-model="editRackDatacenterForm.positionX" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="Z坐标" :label-width="formLabelWidth">
                        <el-input v-model="editRackDatacenterForm.positionZ" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="editRackDatacenterDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="editRackDatacenter">确 定</el-button>
                </div>
            </el-dialog>

            <el-dialog title="新建机柜" :visible.sync="createRackDialogVisible">
                <el-form :model="createRackForm">
                    <el-form-item label="机柜名" :label-width="formLabelWidth">
                        <el-input v-model="createRackForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="总U数" :label-width="formLabelWidth">
                        <el-input v-model="createRackForm.size" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createRackDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createRack">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="物理服务器">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createServerDeviceDialogVisible = true">新建物理服务器</el-button>
            <el-table
                :data="serverDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="hostname"
                    label="主机名">
                </el-table-column>
                <el-table-column
                    fixed
                    prop="ip_info"
                    label="IP信息">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="disk_capacity"
                    label="磁盘(TB)">
                </el-table-column>
                <el-table-column
                    prop="memory_capacity"
                    label="内存(GB)">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="os"
                    label="操作系统">
                </el-table-column>
                <el-table-column
                    prop="rack"
                    label="机柜位置">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="300">
                    <template slot-scope="scope">
                        <el-button @click="showEditDeviceIPDialog(scope.row, 'SERVER_DEVICE')" type="primary" plain size="small">编辑IP信息</el-button>
                        <el-button @click="showEditDeviceRackDialog(scope.row, 'SERVER_DEVICE')" type="primary" plain size="small">编辑机柜位置</el-button>
                        <el-button @click="removeServerDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建物理服务器" :visible.sync="createServerDeviceDialogVisible">
                <el-form :model="createServerDeviceForm">
                    <el-form-item label="主机名" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.hostname" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-select v-model="createServerDeviceForm.brand" placeholder="请选择">
                            <el-option value="IBM">IBM</el-option>
                            <el-option value="DELL">DELL</el-option>
                            <el-option value="惠普">惠普</el-option>
                            <el-option value="华为">华为</el-option>
                            <el-option value="中兴">中兴</el-option>
                            <el-option value="联想">联想</el-option>
                            <el-option value="思科">思科</el-option>
                            <el-option value="Juniper">Juniper</el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="磁盘容量(TB)" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.disk_capacity" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="内存容量(GB)" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.memory_capacity" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createServerDeviceForm.enable_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createServerDeviceForm.expire_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="操作系统" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.os" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createServerDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createServerDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createServerDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="存储设备">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createStorageDeviceDialogVisible = true">新建存储设备</el-button>
            <el-table
                :data="storageDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="name"
                    label="设备名">
                </el-table-column>
                <el-table-column
                    fixed
                    prop="ip_info"
                    label="IP信息">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="rack"
                    label="机柜位置">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="300">
                    <template slot-scope="scope">
                        <el-button @click="showEditDeviceIPDialog(scope.row, 'STORAGE_DEVICE')" type="primary" plain size="small">编辑IP信息</el-button>
                        <el-button @click="showEditDeviceRackDialog(scope.row, 'STORAGE_DEVICE')" type="primary" plain size="small">编辑机柜位置</el-button>
                        <el-button @click="removeStorageDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建存储设备" :visible.sync="createStorageDeviceDialogVisible">
                <el-form :model="createStorageDeviceForm">
                    <el-form-item label="设备名" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-select v-model="createStorageDeviceForm.brand" placeholder="请选择">
                            <el-option value="IBM">IBM</el-option>
                            <el-option value="DELL">DELL</el-option>
                            <el-option value="惠普">惠普</el-option>
                            <el-option value="华为">华为</el-option>
                            <el-option value="中兴">中兴</el-option>
                            <el-option value="联想">联想</el-option>
                            <el-option value="思科">思科</el-option>
                            <el-option value="Juniper">Juniper</el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createStorageDeviceForm.enable_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createStorageDeviceForm.expire_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createStorageDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createStorageDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createStorageDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="网络设备">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createNetworkDeviceDialogVisible = true">新建网络设备</el-button>
            <el-table
                :data="networkDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="name"
                    label="设备名">
                </el-table-column>
                <el-table-column
                    fixed
                    prop="ip_info"
                    label="IP信息">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="rack"
                    label="机柜位置">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="300">
                    <template slot-scope="scope">
                        <el-button @click="showEditDeviceIPDialog(scope.row, 'NETWORK_DEVICE')" type="primary" plain size="small">编辑IP信息</el-button>
                        <el-button @click="showEditDeviceRackDialog(scope.row, 'NETWORK_DEVICE')" type="primary" plain size="small">编辑机柜位置</el-button>
                        <el-button @click="removeNetworkDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建网络设备" :visible.sync="createNetworkDeviceDialogVisible">
                <el-form :model="createNetworkDeviceForm">
                    <el-form-item label="设备名" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-select v-model="createNetworkDeviceForm.brand" placeholder="请选择">
                            <el-option value="IBM">IBM</el-option>
                            <el-option value="DELL">DELL</el-option>
                            <el-option value="惠普">惠普</el-option>
                            <el-option value="华为">华为</el-option>
                            <el-option value="中兴">中兴</el-option>
                            <el-option value="联想">联想</el-option>
                            <el-option value="思科">思科</el-option>
                            <el-option value="Juniper">Juniper</el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createNetworkDeviceForm.enable_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createNetworkDeviceForm.expire_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createNetworkDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createNetworkDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createNetworkDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>

        <el-tab-pane label="其它设备">
            <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="createCommonDeviceDialogVisible = true">新建其它设备</el-button>
            <el-table
                :data="commonDevices"
                border
                highlight-current-row
                style="width: 100%">
                <el-table-column
                    fixed
                    prop="name"
                    label="设备名">
                </el-table-column>
                <el-table-column
                    fixed
                    prop="ip_info"
                    label="IP信息">
                </el-table-column>
                <el-table-column
                    prop="brand"
                    label="厂商">
                </el-table-column>
                <el-table-column
                    prop="model"
                    label="型号">
                </el-table-column>
                <el-table-column
                    prop="enable_time"
                    label="启用时间">
                </el-table-column>
                <el-table-column
                    prop="expire_time"
                    label="过保时间">
                </el-table-column>
                <el-table-column
                    prop="rack"
                    label="机柜位置">
                </el-table-column>
                <el-table-column
                    prop="comment"
                    label="备注">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="300">
                    <template slot-scope="scope">
                        <el-button @click="showEditDeviceIPDialog(scope.row, 'COMMON_DEVICE')" type="primary" plain size="small">编辑IP信息</el-button>
                        <el-button @click="showEditDeviceRackDialog(scope.row, 'COMMON_DEVICE')" type="primary" plain size="small">编辑机柜位置</el-button>
                        <el-button @click="removeCommonDevice(scope.row)" type="danger" plain size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>


            <el-dialog title="新建其它设备" :visible.sync="createCommonDeviceDialogVisible">
                <el-form :model="createCommonDeviceForm">
                    <el-form-item label="设备名" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="厂商" :label-width="formLabelWidth">
                        <el-select v-model="createCommonDeviceForm.brand" placeholder="请选择">
                            <el-option value="IBM">IBM</el-option>
                            <el-option value="DELL">DELL</el-option>
                            <el-option value="惠普">惠普</el-option>
                            <el-option value="华为">华为</el-option>
                            <el-option value="中兴">中兴</el-option>
                            <el-option value="联想">联想</el-option>
                            <el-option value="思科">思科</el-option>
                            <el-option value="Juniper">Juniper</el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="型号" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.model" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="启用时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createCommonDeviceForm.enable_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="过保时间" :label-width="formLabelWidth">
                        <el-date-picker
                            v-model="createCommonDeviceForm.expire_time"
                            type="date"
                            value-format="yyyy-MM-dd"
                            placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                    <el-form-item label="备注" :label-width="formLabelWidth">
                        <el-input v-model="createCommonDeviceForm.comment" autocomplete="off"></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="createCommonDeviceDialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="createCommonDevice">确 定</el-button>
                </div>
            </el-dialog>
        </el-tab-pane>
    </el-tabs>

    
    <el-dialog title="编辑IP信息" :visible.sync="showEditDeviceIPDialogVisible">
        <el-button type="primary" plain size="small" style="float: right; margin-bottom: 5px;" @click="addIPDialogVisible = true">添加IP</el-button>
        <el-table
            :data="deviceIPs"
            border
            highlight-current-row
            style="width: 100%">
            <el-table-column
                prop="ip_address"
                label="IP">
            </el-table-column>
            <el-table-column
                prop="role"
                label="类型">
            </el-table-column>
            <el-table-column
                fixed="right"
                label="操作"
                width="100">
                <template slot-scope="scope">
                    <el-button @click="removeDeviceIP(scope.row)" type="danger" plain size="small">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-dialog>

    <el-dialog title="新增IP" :visible.sync="addIPDialogVisible">
        <el-form :model="addIPForm">
            <el-form-item label="IP" :label-width="formLabelWidth">
                <el-input v-model="addIPForm.ipAddress" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="类型" :label-width="formLabelWidth">
                <el-select v-model="addIPForm.ipRole" placeholder="请选择">
                    <el-option value="业务">业务</el-option>
                    <el-option value="带外">带外</el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="所属网段" :label-width="formLabelWidth">
                <el-select filterable v-model="addIPForm.ipSetId" placeholder="请选择">
                    <el-option
                        v-for="item in ipSets"
                        :key="item.uuid"
                        :label="item.cidr"
                        :value="item.uuid">
                    </el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addIPDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="doAddIP">确 定</el-button>
        </div>
    </el-dialog>

    <el-dialog title="编辑设备位置" :visible.sync="showEditDeviceRackDialogVisible">
        <el-form :model="editDeviceRackForm">
            <el-form-item label="机柜" :label-width="formLabelWidth">
                <el-select filterable v-model="editDeviceRackForm.rackId" placeholder="请选择">
                    <el-option
                        v-for="item in racks"
                        :key="item.uuid"
                        :label="item.name"
                        :value="item.uuid">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="起始U位置(含)" :label-width="formLabelWidth">
                <el-input v-model="editDeviceRackForm.begPos" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="结束U位置(不含)" :label-width="formLabelWidth">
                <el-input v-model="editDeviceRackForm.endPos" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="showEditDeviceRackDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="editDeviceRack">确 定</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios"
import Config from '../../config'

export default {
  name: 'Device',
  data () {
      return {
          config: new Config(),
          formLabelWidth: '120px',

          ipSets: [],

          addIPDialogVisible: false,
          addIPForm: {
              ipAddress: "",
              ipRole: "业务",
              ipSetId: "",
          },
          showEditDeviceIPDialogVisible: false,
          editDeviceIPUUID: "",
          editDeviceIPDeviceType: "",

          showEditDeviceRackDialogVisible: false,
          editDeviceRackUUID: "",
          editDeviceRackDeviceType: "",
          editDeviceRackForm: {
              rackId: "",
              begPos: 0,
              endPos: 0,
          },

          createDataCenterDialogVisible: false,
          datacenters: [],
          createDataCenterForm: {
              name: "",
          },
          
          createRackDialogVisible: false,
          racks: [],
          createRackForm: {
              name: "",
              size: 0,
          },
          editRackDatacenterDialogVisible: false,
          editRackUUID: "",
          editRackDatacenterForm: {
              datacenterId: "",
              positionX: 0,
              positionZ: 0,
          },

          createServerDeviceDialogVisible: false,
          serverDevices: [],
          createServerDeviceForm: {
              brand: "",
              model: "",
              disk_capacity: 0,
              memory_capacity: 0,
              hostname: "",
              enable_time: "",
              expire_time: "",
              os: "",
              comment: "",
          },

          createStorageDeviceDialogVisible: false,
          storageDevices: [],
          createStorageDeviceForm: {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
          },

          createNetworkDeviceDialogVisible: false,
          networkDevices: [],
          createNetworkDeviceForm: {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
          },

          createCommonDeviceDialogVisible: false,
          commonDevices: [],
          createCommonDeviceForm: {
              brand: "",
              model: "",
              name: "",
              enable_time: "",
              expire_time: "",
              comment: "",
          },
      }
  },
  computed: {
      deviceIPs: function() {
          var ips = []
          switch (this.editDeviceIPDeviceType) {
              case "SERVER_DEVICE":
                  for (var device of this.serverDevices) {
                      if (device.uuid == this.editDeviceIPUUID) {
                          ips = device.ips
                          break
                      }
                  }
                  break;
              case "NETWORK_DEVICE":
                  for (var device of this.networkDevices) {
                      if (device.uuid == this.editDeviceIPUUID) {
                          ips = device.ips
                          break
                      }
                  }
                  break;
              case "STORAGE_DEVICE":
                  for (var device of this.storageDevices) {
                      if (device.uuid == this.editDeviceIPUUID) {
                          ips = device.ips
                          break
                      }
                  }
                  break;
              case "COMMON_DEVICE":
                  for (var device of this.commonDevices) {
                      if (device.uuid == this.editDeviceIPUUID) {
                          ips = device.ips
                          break
                      }
                  }
                  break;
          
              default:
                  break;
          }
          return ips
      }
  },
  created () {

  },
  mounted () {
    var that = this
    that.initData()
  },
  methods: {
    initData () {
        var that = this
        that.$store.commit("setPageLoading", true)

        that.ipSets = []

        that.addIPDialogVisible = false
        that.showEditDeviceIPDialogVisible = false
        that.editDeviceIPUUID = ""
        that.editDeviceIPDeviceType = ""
        that.addIPForm = {
              ipAddress: "",
              ipRole: "业务",
              ipSetId: "",
        }

        that.showEditDeviceRackDialogVisible = false
        that.editDeviceRackUUID = ""
        that.editDeviceRackDeviceType = "",
        that.editDeviceRackForm = {
              rackId: "",
              begPos: 0,
              endPos: 0,
        }

        that.createDataCenterDialogVisible = false
        that.datacenters = []
        that.createDataCenterForm = {
              name: "",
        }

        that.createRackDialogVisible = false
        that.racks = []
        that.editRackDatacenterDialogVisible = false
        that.editRackUUID = ""
        that.editRackDatacenterForm = {
              datacenterId: "",
              positionX: 0,
              positionZ: 0,
        }

        that.createServerDeviceDialogVisible = false
        that.serverDevices = []

        that.createStorageDeviceDialogVisible = false
        that.storageDevices = []

        that.createNetworkDeviceDialogVisible = false
        that.networkDevices = []

        that.createCommonDeviceDialogVisible = false
        that.commonDevices = []
        
        Promise.all([
            that.syncIPSets(),
            that.syncDataCenters(),
            that.syncRacks(),
            that.syncServerDevices(),
            that.syncStorageDevices(),
            that.syncNetworkDevices(),
            that.syncCommonDevices()
        ]).then(values => {
            that.$store.commit("setPageLoading", false)
        }).catch(errors => {
            that.$message({
                type: 'error',
                message: "页面加载异常",
                offset: 200,
            })
            console.error(errors)
            that.$store.commit("setPageLoading", false)
        })
    },
    showEditDeviceRackDialog (device, deviceType) {
        this.editDeviceRackUUID = device.uuid
        this.editDeviceRackDeviceType = deviceType
        this.showEditDeviceRackDialogVisible = true
    },
    showEditDeviceIPDialog (device, deviceType) {
        this.editDeviceIPUUID = device.uuid
        this.editDeviceIPDeviceType = deviceType
        this.showEditDeviceIPDialogVisible = true
    },
    syncIPSets () {
        var that = this
        return axios.post(that.config.getAddress("LIST_IPSETS"))
                    .then(response => {
                        that.ipSets = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.ipSets = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeDeviceIP (ip) {
        var that = this
        axios.post(that.config.getAddress("DELETE_IP"), {uuid: ip.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    doAddIP () {
        var that = this
        axios.post(that.config.getAddress("CREATE_IP"), {ip_address: that.addIPForm.ipAddress, ip_type: that.editDeviceIPDeviceType, ip_role: that.addIPForm.ipRole, target_id: that.editDeviceIPUUID, ip_set_id: that.addIPForm.ipSetId})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    syncDataCenters () {
        var that = this
        return axios.post(that.config.getAddress("LIST_DATACENTERS"))
                    .then(response => {
                        that.datacenters = response.data
                    })
                    .catch(error => {
                        console.error(error)
                        that.datacenters = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeDataCenter (datacenter) {
        var that = this
        axios.post(that.config.getAddress("DELETE_DATACENTER"), {uuid: datacenter.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
                console.error(error)
             })
    },
    createDataCenter () {
        var that = this
        axios.post(that.config.getAddress("CREATE_DATACENTER"), {name: that.createDataCenterForm.name})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    syncRacks () {
        var that = this
        return axios.post(that.config.getAddress("LIST_RACKS"))
                    .then(response => {
                        var racks = response.data
                        for (var rack of racks) {
                            if (rack.position.datacenter_uuid !== "") {
                                rack.datacenter = `${rack.position.datacenter_name} (${rack.position.position_x}, ${rack.position.position_z})`
                            }
                        }
                        that.racks = racks
                    })
                    .catch(error => {
                        console.error(error)
                        that.racks = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    showEditRackDatacenterDialog (rack) {
        this.editRackUUID = rack.uuid
        this.editRackDatacenterDialogVisible = true
    },
    editRackDatacenter() {
        var that = this
        axios.post(that.config.getAddress("MAPPING_RACK_DATACENTER"), {rack_id: that.editRackUUID, datacenter_id: that.editRackDatacenterForm.datacenterId, position_x: parseInt(that.editRackDatacenterForm.positionX), position_z: parseInt(that.editRackDatacenterForm.positionZ)})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    editDeviceRack () {
        var that = this
        axios.post(that.config.getAddress("MAPPING_DEVICE_RACK"), {device_id: that.editDeviceRackUUID, device_type: that.editDeviceRackDeviceType, rack_id: that.editDeviceRackForm.rackId, beg_pos: parseInt(that.editDeviceRackForm.begPos), end_pos: parseInt(that.editDeviceRackForm.endPos)})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    removeRack (rack) {
        var that = this
        axios.post(that.config.getAddress("DELETE_RACK"), {uuid: rack.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    createRack () {
        var that = this
        axios.post(that.config.getAddress("CREATE_RACK"), {name: that.createRackForm.name, size: parseInt(that.createRackForm.size)})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    syncServerDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_SERVERS"))
                    .then(response => {
                        var serverDevices = response.data
                        for (var serverDevice of serverDevices) {
                            if (serverDevice.position.rack_uuid !== "") {
                                serverDevice.rack = `${serverDevice.position.rack_name} (${serverDevice.position.beg_pos}U - ${serverDevice.position.end_pos}U)`
                            }

                            serverDevice.ip_info_items = []
                            for (var ip of serverDevice.ips) {
                                serverDevice.ip_info_items.push(`${ip.ip_address}(${ip.role})`)
                            }
                            serverDevice.ip_info = serverDevice.ip_info_items.join(", ")
                        }
                        that.serverDevices = serverDevices
                    })
                    .catch(error => {
                        console.error(error)
                        that.serverDevices = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeServerDevice (server) {
        var that = this
        axios.post(that.config.getAddress("DELETE_SERVER"), {uuid: server.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    createServerDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_SERVER"), {
            brand: that.createServerDeviceForm.brand,
            model: that.createServerDeviceForm.model,
            disk_capacity: parseInt(that.createServerDeviceForm.disk_capacity),
            memory_capacity: parseInt(that.createServerDeviceForm.memory_capacity),
            hostname: that.createServerDeviceForm.hostname,
            enable_time: that.createServerDeviceForm.enable_time,
            expire_time: that.createServerDeviceForm.expire_time,
            os: that.createServerDeviceForm.os,
            comment: that.createServerDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    syncStorageDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_STORAGE_DEVICES"))
                    .then(response => {
                        var storageDevices = response.data
                        for (var storageDevice of storageDevices) {
                            if (storageDevice.position.rack_uuid !== "") {
                                storageDevice.rack = `${storageDevice.position.rack_name} (${storageDevice.position.beg_pos}U - ${storageDevice.position.end_pos}U)`
                            }

                            storageDevice.ip_info_items = []
                            for (var ip of storageDevice.ips) {
                                storageDevice.ip_info_items.push(`${ip.ip_address}(${ip.role})`)
                            }
                            storageDevice.ip_info = storageDevice.ip_info_items.join(", ")
                        }
                        that.storageDevices = storageDevices
                    })
                    .catch(error => {
                        console.error(error)
                        that.storageDevices = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeStorageDevice (device) {
        var that = this
        axios.post(that.config.getAddress("DELETE_STORAGE_DEVICE"), {uuid: device.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    createStorageDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_STORAGE_DEVICE"), {
            brand: that.createStorageDeviceForm.brand,
            model: that.createStorageDeviceForm.model,
            name: that.createStorageDeviceForm.name,
            enable_time: that.createStorageDeviceForm.enable_time,
            expire_time: that.createStorageDeviceForm.expire_time,
            comment: that.createStorageDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    syncNetworkDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_NETWORK_DEVICES"))
                    .then(response => {
                        var networkDevices = response.data
                        for (var networkDevice of networkDevices) {
                            if (networkDevice.position.rack_uuid !== "") {
                                networkDevice.rack = `${networkDevice.position.rack_name} (${networkDevice.position.beg_pos}U - ${networkDevice.position.end_pos}U)`
                            }

                            networkDevice.ip_info_items = []
                            for (var ip of networkDevice.ips) {
                                networkDevice.ip_info_items.push(`${ip.ip_address}(${ip.role})`)
                            }
                            networkDevice.ip_info = networkDevice.ip_info_items.join(", ")
                        }
                        that.networkDevices = networkDevices
                    })
                    .catch(error => {
                        console.error(error)
                        that.networkDevices = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeNetworkDevice (device) {
        var that = this
        axios.post(that.config.getAddress("DELETE_NETWORK_DEVICE"), {uuid: device.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    createNetworkDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_NETWORK_DEVICE"), {
            brand: that.createNetworkDeviceForm.brand,
            model: that.createNetworkDeviceForm.model,
            name: that.createNetworkDeviceForm.name,
            enable_time: that.createNetworkDeviceForm.enable_time,
            expire_time: that.createNetworkDeviceForm.expire_time,
            comment: that.createNetworkDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    syncCommonDevices () {
        var that = this
        return axios.post(that.config.getAddress("LIST_COMMON_DEVICES"))
                    .then(response => {
                        var commonDevices = response.data
                        for (var commonDevice of commonDevices) {
                            if (commonDevice.position.rack_uuid !== "") {
                                commonDevice.rack = `${commonDevice.position.rack_name} (${commonDevice.position.beg_pos}U - ${commonDevice.position.end_pos}U)`
                            }

                            commonDevice.ip_info_items = []
                            for (var ip of commonDevice.ips) {
                                commonDevice.ip_info_items.push(`${ip.ip_address}(${ip.role})`)
                            }
                            commonDevice.ip_info = commonDevice.ip_info_items.join(", ")
                        }
                        that.commonDevices = commonDevices
                    })
                    .catch(error => {
                        console.error(error)
                        that.commonDevices = []
                        that.$message({
                            type: 'error',
                            message: error.response.data.msg,
                            offset: 200,
                        })
                    })
    },
    removeCommonDevice (device) {
        var that = this
        axios.post(that.config.getAddress("DELETE_COMMON_DEVICE"), {uuid: device.uuid})
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    },
    createCommonDevice () {
        var that = this
        axios.post(that.config.getAddress("CREATE_COMMON_DEVICE"), {
            brand: that.createCommonDeviceForm.brand,
            model: that.createCommonDeviceForm.model,
            name: that.createCommonDeviceForm.name,
            enable_time: that.createCommonDeviceForm.enable_time,
            expire_time: that.createCommonDeviceForm.expire_time,
            comment: that.createCommonDeviceForm.comment,
        })
             .then(response => {
                 that.initData()
             })
             .catch(error => {
                console.error(error)
                that.$message({
                    type: 'error',
                    message: error.response.data.msg,
                    offset: 200,
                })
             })
    }
  }
}

</script>