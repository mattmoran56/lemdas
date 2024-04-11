import { getOptions } from "../../apiHelper";

const getUserProfile = async (userId) => {
	const requestOptions = getOptions("GET");

	const response = await fetch(
		`${process.env.REACT_APP_BASE_API_URL}/user/profile/${userId}`,
		requestOptions,
	);
	if (!response.ok) {
		throw new Error("Failed to get profile");
	}

	return response.json();
};

export default getUserProfile;
