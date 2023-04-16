<template>
    <div class="card mx-auto mt-4" style="width: 26rem">
        <div class="card-header text-white bg-success">
            <h3>登录</h3>
        </div>
        <form class="card-body mx-auto" style="width: 20rem">
            <label for="inputTelephone" class="form-label row">
                电话:
                <input
                    type="number"
                    class="form-control"
                    placeholder="请输入手机号"
                    autocomplete="off"
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
                <small class="text-left text-danger" v-if="badPasswordValidate">
                    密码必须为6-18位的字母与数字的组合
                </small>
                <small class="text-left text-success" v-if="goodPasswordValidate">
                    密码强度验证通过
                </small>
            </label>
            <label class="form-check-label my-3" for="stayLoginCheck">
                <input type="checkbox" class="form-check-input" id="stayLoginCheck" />
                7天内保持登录
            </label>
            <div class="row">
                <button type="button" class="btn btn-success btn-block" @click="login()">
                    登录
                </button>
                <a class="btn btn-outline-success btn-block" href="register">
                    没有账号，去注册一个
                </a>
            </div>
        </form>
    </div>
</template>

<script>
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
        ...mapActions("userModule", { userLogin: "login" }),

        login() {
            const re = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$/
            if (re.test(this.user.password)) {
                this.goodPasswordValidate = true
                this.badPasswordValidate = !this.goodPasswordValidate
            } else {
                this.goodPasswordValidate = false
                this.badPasswordValidate = !this.goodPasswordValidate
            }
            if (this.goodTelephoneValidate && this.goodPasswordValidate) {
                this.userLogin(this.user)
                    .then(async () => {
                        this.$bvToast.toast("现在为您自动跳转至主页", {
                            title: "登录成功",
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
    },
}
</script>
<style lang="scss" scoped></style>
