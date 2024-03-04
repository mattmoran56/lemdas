import { getOptions } from "../../apiHelper";

const getUsersGroups = async () => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/group`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get groups");
	}

	return response.json();
};

export default getUsersGroups;
