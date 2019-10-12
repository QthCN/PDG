<template>
    <div :id="uuid"></div>
</template>

<script>
import axios from "axios"
import G2 from '@antv/g2'

import Config from '../../config'

function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

export default {
  name: 'LinePic',
  props: ['title', 'records'],
  data () {
      return {
        chart: null,
        uuid: uuidv4(),
      }
  },
  created () {
    var that = this
  },
  components: {
  },
  mounted () {
    this.render()
  },
  watch: {
  },
  methods: {
    render () {
      var that = this
      that.chart = new G2.Chart({
        container: that.uuid,
        forceFit: true,
        height: 200,
        padding: [20, 80, 60, 40]
      })


      that.chart.scale('值', {
        min: 0
      })

      var tickCount = 10
      if (that.records.length < 10) {
        tickCount = that.records.length 
      }

      that.chart.scale('时间', {
        tickCount: tickCount
      })

      var data = []
      for (var record of this.records) {
        data.push({
          // 加个空格，用于做字符串处理
          "时间": " " + record.key,
          "值": parseFloat(record.value.toFixed(2)),
        })
      }
      that.chart.source(data)

      that.chart.line().position("时间*值")

      that.chart.point().position('时间*值').size(4).shape('circle').style({
        stroke: '#fff',
        lineWidth: 1
      });

      that.chart.render()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
