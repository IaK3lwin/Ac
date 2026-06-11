
ac() {
	local result

	result="$("/home/ian-kelwin/Documentos/projetos/ac/ac" "$@")"

	if [[ "$result" == __AC_CD__:* ]]; then
		cd "${result#__AC_CD__:}"
		return
	fi

	echo "$result"
}
