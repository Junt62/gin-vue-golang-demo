import storageService from "@/service/storageService"
import userService from "@/service/userService"

// function sleep(time) {
//     /* eslint no-promise-executor-return: "off" */
//     return new Promise((resolve) => setTimeout(resolve, time))
// }

const userModule = {
    namespaced: true,
    state: {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: storageService.get(storageService.USER_INFO)
            ? JSON.parse(storageService.get(storageService.USER_INFO))
            : null,
    },
    mutations: {
        SET_TOKEN(state, token) {
            // 更新本地缓存
            storageService.set(storageService.USER_TOKEN, token)
            // 更新state
            state.token = token
        },
        SET_USERINFO(state, userInfo) {
            // 更新本地缓存
            storageService.set(storageService.USER_INFO, JSON.stringify(userInfo))
            // 更新state
            state.userInfo = userInfo
        },
    },
    actions: {
        register(context, { name, telephone, password }) {
            return new Promise((resolve, reject) => {
                userService
                    .register({ name, telephone, password })
                    .then((res) => {
                        // 保存token
                        context.commit("SET_TOKEN", res.data.data.token)
                        // 再次被上一行取代
                        // this.$store.commit("userModule/SET_TOKEN", res.data.data.token)
                        // 被上一行取代
                        // storageService.set(storageService.USER_TOKEN, res.data.data.token)
                        //
                        // 从这里开始使用了链式调用
                        return userService.info()
                    })
                    .then(async (res) => {
                        // 保存用户信息
                        context.commit("SET_USERINFO", res.data.data.user)
                        // 再次被上一行取代
                        // this.$store.commit("userModule/SET_USERINFO", response.data.data.user)
                        // 被上一行取代
                        // storageService.set(
                        //     storageService.USER_INFO,
                        //     JSON.stringify(response.data.data.user),
                        // )
                        // 跳转主页
                        // this.$bvToast.toast("现在为您自动登录", {
                        //     title: "注册成功",
                        //     variant: "success",
                        //     solid: true,
                        // })
                        // await sleep(3000)
                        // this.$router.replace({ name: "home" })
                        resolve(res)
                    })
                    .catch((err) => {
                        reject(err)
                    })
            })
        },
        login(context, { telephone, password }) {
            return new Promise((resolve, reject) => {
                userService
                    .login({ telephone, password })
                    .then((res) => {
                        // 保存token
                        context.commit("SET_TOKEN", res.data.data.token)
                        return userService.info()
                    })
                    .then((res) => {
                        context.commit("SET_USERINFO", res.data.data.user)
                        resolve(res)
                    })
                    .catch((err) => {
                        reject(err)
                    })
            })
        },
        logout({ commit }) {
            // 清除token
            commit("SET_TOKEN", "")
            storageService.set(storageService.USER_TOKEN, "")
            // 清除用户信息
            commit("SET_USERINFO", "")
            storageService.set(storageService.USER_INFO, "")

            window.location.reload()
        },
    },
}

export default userModule
