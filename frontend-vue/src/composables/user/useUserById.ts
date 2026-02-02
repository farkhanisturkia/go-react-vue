// import hook useQuery from vue-query
import { useQuery } from '@tanstack/vue-query';

//import service Api
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

//interface User
export interface User {
    id: number;
    name: string;
    username: string;
    email: string;
}

//composable useUserById dengan parameter id dan return type User
export const useUserById = (id: number) => {

    return useQuery<User, Error>({

        //query key, disesuaikan dengan ID user untuk caching
        queryKey: ['user', id],

        //query function
        queryFn: async () => {

            //get token from cookies
            const token = Cookies.get('token');

            //get user by id from api
            const response = await Api.get(`/api/users/${id}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            //return data
            return response.data.data as User;
        },

    });
};
