<template>
    <div class="register card mx-auto mt-4" style="width: 26rem">
        <div class="card-header text-white bg-primary">
            <h3>注册</h3>
        </div>
        <form class="card-body mx-auto" style="width: 20rem">
            <label for="inputName" class="form-label row">
                姓名:
                <input
                    type="text"
                    class="form-control"
                    placeholder="可选，留空会生成随机姓名"
                    v-model="user.name" />
            </label>
            <label for="inputTelephone" class="form-label row">
                电话:
                <input
                    type="number"
                    class="form-control"
                    placeholder="请输入手机号"
                    v-model="user.telephone" />
                <small class="text-left text-danger" v-if="badTelephoneValidate">
                    手机号必须为1开头的11位数字
                </small>
                <small class="text-left text-success" v-if="goodTelephoneValidate">
                    手机号格式验证成功
                </small>
            </label>
            <label for="inputPassword" class="form-label row">
                密码:
                <input
                    type="password"
                    class="form-control"
                    placeholder="请输入密码"
                    autocomplete="off"
                    v-model="user.password" />
                <small class="text-left text-danger" v-if="badPasswordValidate"
                    >密码必须为6-18位的字母与数字的组合
                </small>
                <small class="text-left text-success" v-if="goodPasswordValidate"
                    >密码强度验证通过
                </small>
            </label>
            <label class="form-check-label my-3" for="autoLoginCheck">
                <input type="checkbox" class="form-check-input" id="autoLoginCheck" />
                自动登录
            </label>
            <div class="row">
                <button type="button" class="btn btn-primary btn-block" @click="register()">
                    注册
                </button>
                <a class="btn btn-outline-primary btn-block" href="login">已有账号，去登录</a>
            </div>
        </form>
    </div>
</template>

<script>
// import userService from "@/service/userService"
// import storageService from "@/service/storageService"
import { mapActions } from "vuex"

function sleep(time) {
    /* eslint no-promise-executor-return: "off" */
    return new Promise((resolve) => setTimeout(resolve, time))
}

export default {
    data() {
        return {
            user: {
                name: "",
                telephone: "",
                password: "",
            },
            goodTelephoneValidate: null,
            badTelephoneValidate: null,
            goodPasswordValidate: null,
            badPasswordValidate: null,
        }
    },
    methods: {
        ...mapActions("userModule", { userRegister: "register" }),
        register() {
            if (this.goodTelephoneValidate && this.goodPasswordValidate) {
                // const api = "http://localhost:1016/api/auth/register"
                // this.$http
                //     .post(api, { ...this.user })
                //     .then((res) => {
                //         // 保存token
                //         console.log(res.data)
                //         // 跳转主页
                //     })
                //     .catch((err) => {
                //         console.log("err:", err.response.data.msg)
                //         this.$bvToast.toast(err.response.data.msg, {
                //             title: "错误",
                //             variant: "danger",
                //             solid: true,
                //         })
                //     })
                this.userRegister(this.user)
                    .then(async () => {
                        this.$bvToast.toast("现在为您自动登录", {
                            title: "注册成功",
                            variant: "success",
                            solid: true,
                        })
                        await sleep(3000)
                        this.$router.replace({ name: "home" })
                    })
                    .catch((err) => {
                        this.$bvToast.toast(err.response.data.msg, {
                            title: "错误",
                            variant: "danger",
                            solid: true,
                        })
                    })
                // userService
                // .register(this.user)
                // .then((res) => {
                // 保存token
                // this.SET_TOKEN(res.data.data.token)
                // 再次被上一行取代
                // this.$store.commit("userModule/SET_TOKEN", res.data.data.token)
                // 被上一行取代
                // storageService.set(storageService.USER_TOKEN, res.data.data.token)
                //
                // 从这里开始使用了链式调用
                // return userService.info()
                // })
                // .then(async (response) => {
                // 保存用户信息
                // this.SET_USERINFO(response.data.data.user)
                // 再次被上一行取代
                // this.$store.commit("userModule/SET_USERINFO", response.data.data.user)
                // 被上一行取代
                // storageService.set(
                //     storageService.USER_INFO,
                //     JSON.stringify(response.data.data.user),
                // )
                // 跳转主页
                //     this.$bvToast.toast("现在为您自动登录", {
                //         title: "注册成功",
                //         variant: "success",
                //         solid: true,
                //     })
                //     await sleep(3000)
                //     this.$router.replace({ name: "home" })
                // })
            } else {
                this.$bvToast.toast("请检查输入的信息", {
                    title: "错误",
                    variant: "danger",
                    solid: false,
                })
                if (this.badTelephoneValidate == null && this.badPasswordValidate == null) {
                    this.badTelephoneValidate = true
                    this.badPasswordValidate = true
                }
            }
        },
    },
    watch: {
        "user.telephone": {
            handler() {
                const re = /^1\d{10}$/
                if (re.test(this.user.telephone)) {
                    this.goodTelephoneValidate = true
                    this.badTelephoneValidate = !this.goodTelephoneValidate
                } else {
                    this.goodTelephoneValidate = false
                    this.badTelephoneValidate = !this.goodTelephoneValidate
                }
            },
        },
        "user.password": {
            handler() {
                const re = /^[A-Za-z0-9]{6,18}$/
                if (re.test(this.user.password)) {
                    this.goodPasswordValidate = true
                    this.badPasswordValidate = !this.goodPasswordValidate
                } else {
                    this.goodPasswordValidate = false
                    this.badPasswordValidate = !this.goodPasswordValidate
                }
            },
        },
    },
}
</script>
<style lang="scss" scoped></style>
