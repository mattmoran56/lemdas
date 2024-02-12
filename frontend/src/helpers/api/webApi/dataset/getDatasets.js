import { getOptions } from "../../apiHelper";

const GetDatasets = async (orderBy) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/dataset${orderBy ? `?orderBy=${orderBy}` : ""}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get datasets");
	}

	return response.json();
};

export default GetDatasets;
