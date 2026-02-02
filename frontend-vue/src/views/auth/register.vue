<script setup lang="ts">

//import ref dan reactive dari vue
import { ref, reactive } from "vue";

//import useRouter dari vue-router
import { useRouter } from "vue-router";

//import composable useRegister
import { useRegister } from "../../composables/auth/useRegister";

//interface untuk error
interface ValidationErrors {
    [key: string]: string;
}

//inisialisasi router
const router = useRouter();

// form state
const name = ref<string>('');
const username = ref<string>('');
const email = ref<string>('');
const password = ref<string>('');

// error state
const errors = reactive<ValidationErrors>({});

// custom composable
const { mutate, isPending } = useRegister();

const handleRegister = (e: Event) => {
    e.preventDefault();

    //call mutate function from useRegister
    mutate(
        {
            name: name.value,
            username: username.value,
            email: email.value,
            password: password.value,
        },
        {
            onSuccess: () => {

                //redirect to login page after success
                router.push("/login");
            },
            onError: (error: any) => {

                //assign errors with the error response
                Object.assign(errors, error.response.data.errors);
            },
        }
    );
};
</script>

<template>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <div class="card border-0 rounded-4 shadow-sm">
                <div class="card-body">
                    <h4 class="fw-bold text-center">REGISTER</h4>
                    <hr />
                    <form @submit.prevent="handleRegister">
                        <div class="row">
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Full Name</label>
                                    <input v-model="name" type="text" class="form-control" placeholder="Full Name" />
                                    <div v-if="errors.Name" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Name }}
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Username</label>
                                    <input v-model="username" type="text" class="form-control" placeholder="Username" />
                                    <div v-if="errors.Username" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Username }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Email address</label>
                                    <input v-model="email" type="email" class="form-control"
                                        placeholder="Email Address" />
                                    <div v-if="errors.Email" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Email }}
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Password</label>
                                    <input v-model="password" type="password" class="form-control"
                                        placeholder="Password" />
                                    <div v-if="errors.Password" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Password }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <button type="submit" class="btn btn-primary w-100 rounded-4" :disabled="isPending">
                            {{ isPending ? 'Loading...' : 'REGISTER' }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>
