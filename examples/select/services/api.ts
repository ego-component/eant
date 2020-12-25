import request from "@/utils/request";

export default {
    select: async () => {
        return request(`/api/select`, {
            method: "GET",
        });
    }
}