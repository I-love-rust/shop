import type { ProductDto } from "$lib/dto/product";
import { restController } from "./controller";

export type createProductParams = {
  name: string;
  description: string;
  tags: string[];
  preview: string;
  price: number;
  characteristics: string;
}

export type createProductResponse = {
  path: string;
}

export type getProductParams = {
  path: string;
}

export type getProductFeedParams = {
  id: number;
}

export type searchProductsParams = {
  page: number | null;
  tags: string[] | string | null;
  text: string | null;
  authorID: number | null;
}

export async function createProduct(params: createProductParams)  {
  return restController.authCall<createProductResponse>({
    path: "/product/new",
    method: "POST",
    body: params
  })
}

export async function getProduct(params: getProductParams): Promise<ProductDto> {
  return restController.call<ProductDto>({
    path: `/product/${params.path}`,
    method: "GET"
  })
}

export async function getProductFeed(params: getProductFeedParams): Promise<ProductDto[]> {
  return restController.call<ProductDto[]>({
    path: `/feed_content?page=${params.id}`,
    method: "GET"
  })
}

export async function searchProduct(params: searchProductsParams): Promise<ProductDto[]> {
  return restController.call<ProductDto[]>({
    path: `/search_product?page=${params.page}&tags[]=${params.tags?.toString() ? params.tags?.toString():"" }&text=${params.text ? params.text : ""}`,
    method: "GET"
  })
}