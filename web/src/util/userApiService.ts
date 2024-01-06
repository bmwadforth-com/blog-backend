import BaseApiService from "./baseApiService";
import {IApiResponse} from "../store/base";

export interface UserStatusResponse {
    userName: string;
    token: string;
    loggedInSince: string;
}

export default class UserApiService extends BaseApiService {

    public async loginUser(username: string, password: string): Promise<IApiResponse<string>> {
        const res = await this.client.post<IApiResponse<string>>('/login', {username, password}, {
            withCredentials: true
        });

        window.localStorage.setItem('token', res.data.data as string);
        
        return res.data;
    }

    public async userStatus(): Promise<IApiResponse<UserStatusResponse>> {
        const res = await this.client.get<IApiResponse<UserStatusResponse>>('/user/status', {
            withCredentials: true
        });

        return res.data;
    }
    
}