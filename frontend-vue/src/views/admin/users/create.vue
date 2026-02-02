<script setup lang="ts">

//import ref dan reactive dari vue
import { ref, reactive } from 'vue'

//import useRouter dari vue-router
import { useRouter } from 'vue-router'

//import component SidebarMenu
import SidebarMenu from '../../../components/SidebarMenu.vue'

//import composable useUserCreate
import { useUserCreate } from '../../../composables/user/useUserCreate'

// Router instance
const router = useRouter()

// Form state
const name = ref<string>('')
const username = ref<string>('')
const email = ref<string>('')
const password = ref<string>('')

// Validation errors
const errors = reactive<Record<string, string>>({})

// Mutation
const { mutate, isPending } = useUserCreate()

// Handle form submit
const storeUser = (e: Event) => {
    e.preventDefault()

    // Call user create mutation
    mutate(
        {
            name: name.value,
            username: username.value,
            email: email.value,
            password: password.value,
        },
        {
            onSuccess: () => {

                // Redirect to users index
                router.push('/admin/users')

            },
            onError: (error: any) => {

                // Assign validation errors to the errors object
                Object.assign(errors, error.response.data.errors)

            },
        }
    )
}
</script>

<template>
    <div class="container mt-5 mb-5">
        <div class="row">
            <div class="col-md-3">
                <SidebarMenu />
            </div>
            <div class="col-md-9">
                <div class="card border-0 rounded-4 shadow-sm">
                    <div class="card-header">
                        ADD USER
                    </div>
                    <div class="card-body">
                        <form @submit="storeUser">
                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Full Name</label>
                                <input type="text" v-model="name" class="form-control" placeholder="Full Name" />
                                <div v-if="errors.Name" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.Name }}
                                </div>
                            </div>

                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Username</label>
                                <input type="text" v-model="username" class="form-control" placeholder="Username" />
                                <div v-if="errors.Username" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.Username }}
                                </div>
                            </div>

                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Email address</label>
                                <input type="email" v-model="email" class="form-control" placeholder="Email Address" />
                                <div v-if="errors.Email" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.Email }}
                                </div>
                            </div>

                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Password</label>
                                <input type="password" v-model="password" class="form-control" placeholder="Password" />
                                <div v-if="errors.Password" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.Password }}
                                </div>
                            </div>

                            <button type="submit" class="btn btn-md btn-primary rounded-4 shadow-sm border-0"
                                :disabled="isPending">
                                {{ isPending ? 'Saving...' : 'Save' }}
                            </button>

                            <router-link to="/admin/users"
                                class="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2">
                                Cancel
                            </router-link>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
