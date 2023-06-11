import           Control.Monad
import qualified Data.Char     as Char
import qualified Data.Set      as Set

testData = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"

makeHalf :: String -> (String, String)
makeHalf str = (take h str, drop h str)
  where
    h =  (flip div) 2 $ length str

intersectHalfs :: (String, String) -> Set.Set Char
intersectHalfs (a, b) = Set.intersection sa sb
  where
    sa = Set.fromList a
    sb = Set.fromList b

intersectTriad :: [String] -> Set.Set Char
intersectTriad (x:y:z:_) =  Set.intersection sz $ Set.intersection sx sy
  where
    sx = Set.fromList x
    sy = Set.fromList y
    sz = Set.fromList z


findPriority :: Char -> Int
findPriority c
  | c `elem` ['A'..'Z'] = Char.ord c  - Char.ord 'A' + 27
  | c `elem` ['a'..'z'] = Char.ord c  - Char.ord 'a' + 1
  | otherwise = 0

partition :: Int -> [a] -> [[a]]
partition _ [] = []
partition n xs = (take n xs) : (partition n (drop n xs))

solve1st = sum . map (findPriority . Set.elemAt 0 . intersectHalfs . makeHalf)
solve2nd = sum . map (findPriority . Set.elemAt 0 .intersectTriad) . partition 3

main = do
    let testInput = lines testData
    unless ((solve1st testInput) == 157) (error "test fail for part one")
    unless ((solve2nd testInput) == 70)  (error "test fail for part two")

    contents <- readFile "dat/day03.txt"
    let input = lines contents
    --
    putStrLn $ show $ solve1st input
    putStrLn $ show $ solve2nd input
