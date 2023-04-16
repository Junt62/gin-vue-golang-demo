const userRoute = [
    {
        path: "/register",
        name: "register",
        component: () => import("@/views/register/RegisterView.vue"),
    },
    {
        path: "/login",
        name: "login",
        component: () => import("@/views/login/LoginView.vue"),
    },
    {
        path: "/profile",
        name: "profile",
        meta: {
            auth: true,
        },
        component: () => import("@/views/profile/ProfileView.vue"),
    },
]

export default userRoute
