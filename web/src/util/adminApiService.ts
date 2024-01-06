import BaseApiService from "./baseApiService";
import {IApiResponse} from "../store/base";
export default class AdminApiService extends BaseApiService {
    public async queryGemini(query: string): Promise<string> {
        const res = await this.client.get<string>('/gemini', {
            withCredentials: true,
            params: {
                query
            }
        });

        return res.data;
    }
}