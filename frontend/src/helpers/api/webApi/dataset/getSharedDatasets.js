import { getOptions } from "../../apiHelper";

const getSharedDatasets = async () => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/datasets/shared`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get datasets");
	}
	const data = await response.json();

	return data.datasets;
};

export default getSharedDatasets;
