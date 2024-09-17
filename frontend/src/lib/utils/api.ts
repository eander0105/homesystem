const BASE_URL = 'http://localhost:8080';

type QueryParams = Record<string, string | number | boolean | undefined>;

function handleQueryParams(queryParams: QueryParams = {}): string {
    const params = Object.keys(queryParams)
        .filter(key => queryParams[key] !== undefined && queryParams[key] !== null) // filter out undefined or null values
        .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(String(queryParams[key]))}`) // safely encode
        .join('&');

    return params ? `?${params}` : ''; // only add '?' if there are query params
}

async function request(method: string, endpoint: string, data?: any, queryParams?: QueryParams) {
    const url = `${BASE_URL}${endpoint}${handleQueryParams(queryParams)}`;

    const options: RequestInit = {
        method,
        headers: {
            'Content-Type': 'application/json',
        },
    };

    if (data) {
        options.body = JSON.stringify(data); // Only add body if data exists
    }

    try {
        const response = await fetch(url, options);
        return await handleResponse(response);
    } catch (error: any) {
        // Enhanced error handling, possibly log or throw a more descriptive error
        console.error('API request error:', error);
        throw new Error(`API request failed: ${error.message}`);
    }
}

async function get<T = any>(endpoint: string, queryParams?: QueryParams): Promise<T> {
    console.log('GET', endpoint, queryParams);
    return request('GET', endpoint, undefined, queryParams);
}

async function post<T = any>(endpoint: string, data: Partial<T>): Promise<T> {
    return request('POST', endpoint, data);
}

async function put<T = any>(endpoint: string, data: Partial<T>): Promise<T> {
    return request('PUT', endpoint, data);
}

async function del<T = any>(endpoint: string): Promise<T> {
    return request('DELETE', endpoint);
}

async function handleResponse(response: Response): Promise<any> {
    if (!response.ok) {
        const error = await response.text();
        // You can add more granular error handling based on response status codes
        throw new Error(`Error: ${response.status} - ${error}`);
    }
    return response.json();
}

export default { get, post, put, del };
