<template>
  <el-menu
    mode="horizontal"
    :router="true"
  >
    <el-menu-item
      index="0"
      :route="{ name: 'Explore' }"
    >
      知识广场
    </el-menu-item>
    <el-menu-item
      index="1"
      :route="{ name: 'Answerer' }"
    >
      我要提问
    </el-menu-item>
    <el-menu-item
      index="2"
      :route="{ name: 'Question' }"
    >
      我的问题
    </el-menu-item>
    <el-sub-menu index="3">
      <template #title>
        个人信息
      </template>
      <el-menu-item
        index="4"
        :route="{ name: 'User' }"
      >
        用户资料
      </el-menu-item>
      <el-menu-item
        index="5"
        :route="{ name: 'Income' }"
      >
        收入统计
      </el-menu-item>
      <el-menu-item
        index="6"
        @click="logout"
      >
        退出登录
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script>
export default {
  name: 'NavBar',
  methods: {
    logout() {
      window.localStorage.removeItem('token');
      this.$message({
        message: '登出成功',
        type: 'success',
      });
      this.$router.push({
        name: 'Explore',
      });
    },
  },
};
</script>

<style scoped></style>
