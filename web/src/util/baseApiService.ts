import {AxiosInstance} from "axios";
import axios from 'axios';
import {getToken, getTokenString} from "./tokenUtil";

export default abstract class BaseApiService {
    protected readonly client: AxiosInstance;
    constructor() {
        const axiosClient = axios.create({
            baseURL: process.env.REACT_APP_API_URL
        });

        // Add a request interceptor
        axiosClient.interceptors.request.use(function (config) {
            const token = getTokenString();
            if (token !== null){
                // Ensuring we're not adding authorization header when fetching from CDN.
                if (config.baseURL === process.env.REACT_APP_API_URL) {
                    config.headers.Authorization = `Bearer ${token}`;
                }
            }

            return config;
        }, function (error) {
            // Do something with request error
            return Promise.reject(error);
        });

        // Add a response interceptor
        axiosClient.interceptors.response.use(function (response) {
            // Any status code that lie within the range of 2xx cause this function to trigger
            // Do something with response data
            return response;
        }, function (error) {
            // Any status codes that falls outside the range of 2xx cause this function to trigger
            // Do something with response error
            return Promise.reject(error);
        });

        this.client = axiosClient;
    }
}