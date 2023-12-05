let rec step_one ic accum =
    let line = In_channel.input_line ic in
    match line with
    | Some line ->
        let space_split = String.split_on_char ' ' line in

    | None -> accum


let () =
    let ic = open_in "./test.txt" in
    let res = step_one ic [] in
