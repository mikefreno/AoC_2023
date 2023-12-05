(*
   take in each line
   get the first number in line
   reverse line, get first number
   put them together as strings
   parse to int
   add to running total
*)

let is_digit c =
  let code = Char.code c in
  code >= 48 && code <= 57
;;

let is_spelt = function
  | [ 'o'; 'n'; 'e' ] -> '1'
  | [ 't'; 'w'; 'o' ] -> '2'
  | [ 't'; 'h'; 'r'; 'e'; 'e' ] -> '3'
  | [ 'f'; 'o'; 'u'; 'r' ] -> '4'
  | [ 'f'; 'i'; 'v'; 'e' ] -> '5'
  | [ 's'; 'i'; 'x' ] -> '6'
  | [ 's'; 'e'; 'v'; 'e'; 'n' ] -> '7'
  | [ 'e'; 'i'; 'g'; 'h'; 't' ] -> '8'
  | [ 'n'; 'i'; 'n'; 'e' ] -> '9'
  | _ -> '0'
;;

(*let rec get_first_num = function*)
(*| hd :: _ when is_digit hd -> hd;*)
(*| _ :: tl ->  get_first_num tl*)
(*| [] -> failwith ("no num found")*)
(*;;*)

(*let rec rev_helper list accum =*)
(*match list with*)
(*| [] -> accum*)
(*| hd :: tl -> rev_helper tl (hd :: accum)*)
(*;;*)

let rec cycler accum =
  (*Quick_print.chars_list accum;*)
  match accum with
  | x when is_spelt x != '0' -> is_spelt x
  | _ :: tl -> cycler tl
  | [] -> '0'
;;

let drop_last_entry list =
  match List.rev list with
  | [] -> []
  | _ :: tl -> List.rev tl
;;

let rec rev_cycler accum =
  (*Quick_print.chars_list accum;*)
  match accum with
  | x when is_spelt x != '0' -> is_spelt x
  | _ :: _ as x -> rev_cycler (drop_last_entry x)
  | [] -> '0'
;;

let rec help_me list accum =
  let x = cycler accum in
  match x with
  | '0' ->
    (match list with
     | hd :: _ when is_digit hd -> hd
     | hd :: tl -> help_me tl (accum @ [ hd ])
     | [] -> failwith "found no num")
  | n -> n
;;

let rec help_me_rev list accum =
  let x = rev_cycler accum in
  match x with
  | '0' ->
    (match list with
     | hd :: _ when is_digit hd -> hd
     | hd :: tl -> help_me_rev tl (hd :: accum)
     | [] -> failwith "found no num")
  | n -> n
;;

(*let rec accum_total ic accum =*)
(*let line = In_channel.input_line ic in*)
(*match line with*)
(*| Some line -> *)
(*let chars = List.of_seq (String.to_seq line) in*)
(*let left = get_first_num chars in*)
(*let right = get_first_num (rev_helper chars []) in*)
(*let as_str = ((String.make 1 left) ^ (String.make 1 right)) in*)
(*accum_total ic (accum + (int_of_string as_str));*)
(*| None -> accum*)

let rec accum_total ic accum =
  let line = In_channel.input_line ic in
  match line with
  | Some line ->
    let chars = List.of_seq (String.to_seq line) in
    let left = help_me chars [] in
    (*print_string "left: ";*)
    (*print_char left;*)
    (*print_char '\n';*)
    let right = help_me_rev (List.rev chars) [] in
    (*print_string "right: ";*)
    (*print_char right;*)
    (*print_char '\n';*)
    let as_str = String.make 1 left ^ String.make 1 right in
    accum_total ic (accum + int_of_string as_str)
  | None -> accum
;;

let () =
  let ic = open_in "./input.txt" in
  (*let part_one_res = accum_total ic 0 in*)
  (*print_string "part one: ";*)
  (*print_int part_one_res;*)
  print_int (accum_total ic 0)
;;
