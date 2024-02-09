import { useState, useEffect } from "react";
import { jwtDecode } from "jwt-decode";

const useAuth = () => {
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);
	const [username, setUsername] = useState(null);

	useEffect(() => {
		const token = localStorage.getItem("token");

		if (token) {
			try {
				// TODO: check if expired
				const decodedToken = jwtDecode(token);
				const name = decodedToken.first_name;
				setUsername(name);
			} catch (err) {
				setError("Invalid token");
			}
		}

		setLoading(false);
	}, []);

	return { username, error, loading };
};

export default useAuth;
