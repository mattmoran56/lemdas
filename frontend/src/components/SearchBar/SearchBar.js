import React, { useEffect, useState } from "react";
import { ChevronDownIcon, MagnifyingGlassIcon, ChevronUpIcon } from "@heroicons/react/24/outline";
import { useNavigate, useSearchParams } from "react-router-dom";
import Button from "../basic/Button";
import OptionSection from "./OptionSection";
import useAuth from "../../hooks/useAuth";
import logout from "../../helpers/utils/logout";
import Logo from "../basic/Logo";

const SearchBar = () => {
	const [searchTerm, setSearchTerm] = useState("");
	const [expanded, setExpanded] = useState(false);
	const [options, setOptions] = useState([{ logic: "and", field: "title", term: "" }]);
	const [slideoutTop, setSlideoutTop] = useState(-1000);

	const { username } = useAuth();

	const navigate = useNavigate();

	const [searchParams] = useSearchParams();

	const handleSearch = () => {
		navigate(`/search?query=${searchTerm}`);
	};
	const handleComplexSearch = () => {
		// TODO: handle complex search
	};

	useEffect(() => {
		if (searchParams.get("query") !== "") {
			setSearchTerm(searchParams.get("query"));
		}
	}, [searchParams]);

	useEffect(() => {
		const slideout = document.getElementById("slideout");
		setSlideoutTop(slideout.clientHeight * -1);
	}, [options]);

	return (
		<div
			className="w-full max-h-min transition-all duration-500"
			style={{
				marginBottom: (expanded ? "0" : `${slideoutTop}px`),
			}}
		>
			<div
				className={`w-full bg-oxfordblue text-lightlavender p-4 flex items-center justify-between
							shadow-md shadow-[#555] z-50 relative`}
			>
				<a href="/" aria-label="home"><Logo className="mr-4" /></a>
				<div className="flex flex-auto justify-center">
					<input
						placeholder="Search datasets"
						className="px-3 py-2 max-w-96 w-1/2 rounded-3xl text-black outline-none"
						value={searchTerm}
						onChange={(event) => { return setSearchTerm(event.target.value); }}
						onKeyPress={(e) => { if (e.key === "Enter") { handleSearch(); } }}
					/>
					<Button className="mx-2" onClick={handleSearch}>
						Search
						<MagnifyingGlassIcon className="w-4 h-4 ml-2" />
					</Button>
					<Button
						className="mx-2 bg-transparent border-transparent"
						onClick={() => { return setExpanded(!expanded); }}
					>
						Options
						{expanded
							? <ChevronUpIcon className="w-4 h-4 ml-2" />
							: <ChevronDownIcon className="w-4 h-4 ml-2" /> }
					</Button>
				</div>
				<div>
					{!username ? (
						<Button className="mx-2" onClick={() => { navigate("/login"); }}>
							Login
						</Button>
					) : (
						<div className="flex">
							<Button className="mx-2" onClick={() => { navigate("/datasets"); }}>
								My Data
							</Button>
							<Button
								className="mx-2 bg-offwhite border-offwhite !text-indianred"
								onClick={() => { logout(); }}
							>
								Logout
							</Button>
						</div>
					)}
				</div>
			</div>
			<div
				className={`bg-lightlavender z-10 relative transition-all duration-500 py-2 px-12
							flex justify-center items-center`}
				style={{
					top: (expanded ? "0" : `${slideoutTop}px`),
				}}
				id="slideout"
			>
				<div className="flex justify-center items-center flex-col">
					{options.map((option, index) => {
						return (
							<OptionSection
								key={option.id}
								first={index === 0}
								lastRemaining={options.length === 1}
								logic={option.logic}
								field={option.field}
								term={option.term}
								onChange={(event) => {
									const newOptions = [...options];
									newOptions[index][event.target.id] = event.target.value;
									setOptions(newOptions);
								}}
								onNewOption={() => {
									const newOptions = [...options];
									newOptions.push({ logic: "and", field: "title", term: "" });
									setOptions(newOptions);
								}}
								onDelete={() => {
									const newOptions = [...options];
									newOptions.splice(index, 1);
									setOptions(newOptions);
								}}
							/>
						);
					})}
					<div className="flex justify-end w-full">
						<Button
							className={`mx-2 transition-all duration-500 delay-100
										${options.length === 1 ? "mr-12" : ""}`}
							onClick={handleComplexSearch}
						>
							Search
							<MagnifyingGlassIcon className="w-4 h-4 ml-2" />
						</Button>
					</div>
				</div>
			</div>
		</div>
	);
};

export default SearchBar;
