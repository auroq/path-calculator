module PathCalculator
    ( RoadSystem
    , Section (..)
    , Label (..)
    , Path
    , optimalPath
    , roadStep
    ) where

data Section = Section { getA :: Int, getB :: Int, getC :: Int } deriving (Show, Eq)
type RoadSystem = [Section]

data Label = A | B | C deriving (Show, Eq)
type Path = [(Label, Int)]

optimalPath :: RoadSystem -> Path
optimalPath roadSystem =
    let (bestPathA, bestPathB) = foldl roadStep ([],[]) roadSystem
    in  if sum (map snd bestPathA) <= sum (map snd bestPathB)
            then reverse bestPathA
            else reverse bestPathB

roadStep :: (Path, Path) -> Section -> (Path, Path)
roadStep (pathA, pathB) (Section a b c) =
    let priceA = sum $ map snd pathA
        priceB = sum $ map snd pathB
        forwardPriceToA = priceA + a
        forwardPriceToB = priceB + b
        crossPriceToA = priceB + b + c
        crossPriceToB = priceA + a + c
        newPathToA = if forwardPriceToA <= crossPriceToA
                        then (A,a):pathA
                        else (C,c):(B,b):pathB
        newPathToB = if forwardPriceToB <= crossPriceToB
                        then (B,b):pathB
                        else (C,c):(A,a):pathA
    in  (newPathToA, newPathToB)


