import { useState, useEffect } from "react";
import { jwtDecode } from "jwt-decode";

const useAuth = () => {
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);
	const [username, setUsername] = useState(null);
	const [id, setId] = useState(null);

	useEffect(() => {
		const token = localStorage.getItem("token");

		if (token) {
			try {
				// TODO: check if expired
				const decodedToken = jwtDecode(token);
				const name = decodedToken.first_name;
				const userId = decodedToken.user_id;
				const expired = decodedToken.exp < Date.now() / 1000;
				if (expired) {
					setError("Token expired");
					if (window.location.pathname !== "/login") window.location.href = "/login";
				} else {
					setUsername(name);
					setId(userId);
				}
			} catch (err) {
				setError("Invalid token");
			}
		}

		setLoading(false);
	}, []);

	return {
		username, id, error, loading,
	};
};

export default useAuth;
