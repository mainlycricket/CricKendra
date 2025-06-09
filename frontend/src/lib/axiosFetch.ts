import axios from "axios";

const backendFetch = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL,
  withCredentials: true,
});

export interface IReqConfig<T> {
  url: string;
  method: string;
  responseType?: "arraybuffer" | "blob" | "document" | "json" | "text" | "stream" | "formdata";
  body?: T;
  headers?: {
    [index: string]: string;
  };
}

export async function doBackendRequest<ReqT, ResT>(
  config: IReqConfig<ReqT>
): Promise<IBackendResponse<ResT>> {
  const { url, method, headers, body: data } = config;
  const responseType = config.responseType || "json";

  try {
    const response = await backendFetch.request({ url, method, headers, data, responseType });
    return response.data as IBackendResponse<ResT>;
  } catch (error) {
    console.error(`error while doing request: ${JSON.stringify(error)}`);
    return { success: false, message: "something went wrong..." };
  }
}

export interface IBackendResponse<T> {
  success: boolean;
  message: string;
  data?: T;
}
