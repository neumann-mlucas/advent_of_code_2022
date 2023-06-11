{-# LANGUAGE OverloadedStrings #-}

import           Control.Monad
import qualified Data.List     as List
import qualified Data.Text     as Text

testData = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000\n"

parseInventores :: Text.Text -> [[Int]]
parseInventores inv =
  map parse $ Text.splitOn "\n\n" inv
  where
    txtToInt = (read::String -> Int) . Text.unpack
    parse = map txtToInt . Text.words


solve1st :: [[Int]] -> Int
solve1st = maximum . map sum

solve2nd :: [[Int]] -> Int
solve2nd = sum . List.take 3 . List.reverse . List.sort . map sum

main = do
    let testInput = parseInventores $ Text.pack testData
    unless ((solve1st testInput) == 24000) (error "test fail for part one")
    unless ((solve2nd testInput) == 45000) (error "test fail for part two")

    contents <- readFile "dat/day01.txt"
    let input = parseInventores $ Text.pack contents

    putStrLn $ show $ solve1st input
    putStrLn $ show $ solve2nd input
