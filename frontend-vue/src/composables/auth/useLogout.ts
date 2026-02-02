// import js-cookie
import Cookies from "js-cookie";

// import useRouter from vue-router
import { useRouter } from 'vue-router';

// custom composables useLogout
export const useLogout = (): (() => void) => {

    // initialize router
    const router = useRouter();

    // Fungsi logout
    const logout = (): void => {
        
        // Hapus token dan user dari cookie
        Cookies.remove("token");
        Cookies.remove("user");

        // redirect to login
        router.push({ name: 'login' })
    };

    return logout;
};
