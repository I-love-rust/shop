import { restController } from "./controller";

export type createBillParams = {
    product: number;
  }
  

export async function generateBill(params: createBillParams): Promise<{}> {
    return restController.authCall<{}>({
      path: `/purchase/generate?product=${params.product}`,
      method: "GET"
    })
  }