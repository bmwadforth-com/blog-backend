import {atom, selector} from 'recoil';
import UserApiService from "../../util/userApiService";
import {getToken} from "../../util/tokenUtil";

export interface IUser {
    username: string;
    isLoggedIn: boolean;
    loggedInSince: Date | null;
}

export interface IUSerState extends IUser {}

export const userStateSelector = selector<IUSerState>({
    key: 'userStateSelector',
    get: async ({get}) => {
        const state = get(userState)
        try {
            const tokenData = getToken();
            if (tokenData) {
                const apiService = new UserApiService();
                const data = await apiService.userStatus();
                if (data.data?.active) {
                    return { username:  data.data?.userName as string, loggedInSince: new Date(data.data?.loggedInSince as string) as Date, isLoggedIn: true }
                }
            }

            return state;
        } catch (error) {
            return state;
        }
    },
    set: ({set}, newVal) => set(userState, newVal as IUSerState)
});

const userState = atom<IUSerState>({
    key: 'userState',
    default: { username: '', isLoggedIn: false, loggedInSince: null } as IUSerState
});

export default userState;