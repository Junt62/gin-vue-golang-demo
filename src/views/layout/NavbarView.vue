<template>
    <div id="header">
        <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
            <a class="navbar-brand" href="/">Navbar</a>
            <button
                class="navbar-toggler"
                type="button"
                data-toggle="collapse"
                data-target="#navbarSupportedContent"
                aria-controls="navbarSupportedContent"
                aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav ml-auto">
                    <!-- 用户登录后的下拉菜单 -->
                    <li class="nav-item dropdown" v-if="userInfo">
                        <a
                            class="nav-link dropdown-toggle"
                            href="#"
                            role="button"
                            data-toggle="dropdown"
                            aria-expanded="false">
                            {{ userInfo.name }}
                        </a>
                        <div class="dropdown-menu dropdown-menu-right">
                            <!-- eslint-disable-next-line -->
                            <a class="dropdown-item" href="profile">个人主页</a>
                            <!-- eslint-disable-next-line -->
                            <a class="dropdown-item" @click="logout">登出</a>
                        </div>
                    </li>
                    <!-- 用户登录前显示的登录和注册按钮 -->
                    <li class="nav-item" v-if="!userInfo">
                        <a
                            class="btn btn-outline-success btn-block"
                            href="login"
                            v-if="$route.name != 'login'">
                            登录
                        </a>
                    </li>
                    <li class="nav-item" v-if="!userInfo">
                        <a
                            class="btn btn-primary btn-block"
                            href="register"
                            v-if="$route.name != 'register'">
                            注册
                        </a>
                    </li>
                </ul>
            </div>
        </nav>
    </div>
</template>

<script>
// import storageService from "@/service/storageService"
import { mapState, mapActions } from "vuex"

export default {
    computed: mapState({
        userInfo: (state) => state.userModule.userInfo,
    }),
    // userInfo() {
    //     return this.$store.state.userModule.userInfo
    //     // return JSON.parse(storageService.get(storageService.USER_INFO))
    // },

    methods: mapActions("userModule", ["logout"]),
}
</script>
<style lang="scss" scoped></style>
