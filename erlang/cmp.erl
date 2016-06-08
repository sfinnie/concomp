-module(cmp).
-export([cmp/2]).

cmp(Fname1, Fname2) ->
    Sizes=pfile_size([Fname1, Fname2]),
    Size1 = proplists:get_value(Fname1, Sizes),
    Size2 = proplists:get_value(Fname2, Sizes),
    
    if 
        Size1 > Size2 ->
            io:format("The first file is bigger~n");
        Size2 > Size1 ->
            io:format("The second file is bigger~n");
        true ->
            io:format("The files are the same size~n")
    end.
 
pfile_size(Files) ->
    ReplyTo = self(),
    Keys = [spawn(fun() -> ReplyTo ! {self(), F, filelib:file_size(F)} end) || F <- Files],
    Yield = fun(Key) ->
                receive
                    {Key, Fname, Fsize}   -> {Fname, Fsize}
                end
            end,
    [Yield(Key) || Key <- Keys].

