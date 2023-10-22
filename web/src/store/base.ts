export interface IApiError {
    message: string;
    code: string;
}

export interface IApiResponse<T> {
    message: string;
    data?: T;
    errors?: IApiError[];
}
