// import useMutation dari '@tanstack/vue-query';
import { useMutation } from '@tanstack/vue-query';

//import service API
import Api from '../../services/api';

//interface LoginRequest
interface LoginRequest {
    username: string;
    password: string;
}

export const useLogin = () => {

    return useMutation({

        // mutation untuk login
        mutationFn: async (data: LoginRequest) => {

            //menggunakan service API untuk login
            const response = await Api.post('/api/login', data);

            //mengembalikan response data
            return response.data;
        }
    });
};
