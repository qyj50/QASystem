<template>
  <div>
    <el-container>
      <el-main>
        <el-row justify="center">
          <el-col :span="16">
            <el-alert
              title="扫码支付"
              type="info"
              description="请扫描右方二维码进行支付"
              show-icon
            />
          </el-col>
        </el-row>

        <el-row
          :gutter="2"
          class="paybox"
          justify="center"
        >
          <el-col :span="5">
            <el-image
              style="width: 100; height: 100"
              :src="require('../assets/pay-demo.png')"
              :fit="fill"
            />
          </el-col>
          <el-col :span="3" />
          <el-col :span="5">
            <el-image
              style="width: 100; height: 100"
              :src="require('../assets/pay-qrscan.png')"
              :fit="fill"
            />
          </el-col>
        </el-row>
        <el-row justify="center">
          <el-button
            type="primary"
            @click="onSubmit"
          >
            完成支付
          </el-button>
          <el-button @click="quit">
            取消支付
          </el-button>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'Pay',
  methods: {
    quit() {
      this.$router.go(-1);
    },
    onSubmit() {
      fetch('/v1/question/pay', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('token'),
        },
        body: JSON.stringify({
          questionid: parseInt(this.$route.query.id, 10),
        }),
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('支付失败!');
          }
        })
        .then(() => {
          this.$message({
            message: '支付成功',
            type: 'success',
          });
          this.$router.push({
            name: 'Question',
          });
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
    },
  },
};
</script>

<style scoped>
.el-row {
  margin: 50px auto;
}
.pay-box {
  height: 500px;
}
</style>
