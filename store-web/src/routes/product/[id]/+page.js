// @ts-nocheck
import { getProduct } from '$lib/api/product.ts'

export const load = ({ params }) => {
    const product = async(path) => {
        const res = await getProduct({ path: path })
        return res
    }

    return {
        data: product(params.id)
    }
}