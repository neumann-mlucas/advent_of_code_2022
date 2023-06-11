import           Control.Monad
import qualified Data.Char     as Char
import qualified Data.Set      as Set

testInput = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

uniqueChars :: String -> Int
uniqueChars = Set.size . Set.fromList

getSubStrings :: Int -> String -> [String]
getSubStrings n str
  | length(str) > n = (take n str): (getSubStrings  n (drop 1 str))
  | otherwise = []

getIdx :: Int -> [Bool] -> Int
getIdx _ []         = -1
getIdx n (True:xs)  = n
getIdx n (False:xs) = getIdx (succ n) xs

solve n = (+ n) . (getIdx 0) . map ((== n) . uniqueChars) . getSubStrings n
solve1st = solve 4
solve2nd = solve 14

main = do
    unless ((solve1st testInput) == 7)  (error "test fail for part one")
    unless ((solve2nd testInput) == 19) (error "test fail for part two")

    input <- readFile "dat/day06.txt"

    putStrLn $ show $ solve1st input
    putStrLn $ show $ solve2nd input
