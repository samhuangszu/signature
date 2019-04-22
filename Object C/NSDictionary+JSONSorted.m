#import "NSArray+JSONSorted.h"
#import "NSDictionary+JSONSorted.h"
@implementation NSDictionary (JSONSorted) 
- (NSArray *)sortedKeys
{
    NSSortDescriptor *sortDescriptor = [NSSortDescriptor sortDescriptorWithKey:@"description" ascending:YES selector:@selector(compare:)];
    return [self.allKeys sortedArrayUsingDescriptors:@[sortDescriptor]];
}

- (NSString *)sortedJSONString
{
    NSMutableArray *items = [NSMutableArray array];
    for (NSString *key in [self sortedKeys]) {
        id value = [self objectForKey:key];
        
        if ([value isKindOfClass:[NSDictionary class]]) {
            NSString *itemString = [NSString stringWithFormat:@"\"%@\":%@", key, [value sortedJSONString]];
            [items addObject:itemString];
        }
        else if ([value isKindOfClass:[NSArray class]]) {
            NSString *itemString = [NSString stringWithFormat:@"\"%@\":%@", key, [value sortedJSONString]];
            [items addObject:itemString];
        }
        else {
            if ([value isKindOfClass:[NSString class]]) {
                value = [NSString stringWithFormat:@"\"%@\"", value];
            }
            NSString *itemString = [NSString stringWithFormat:@"\"%@\":%@", key, value];
            [items addObject:itemString];
        }
    }
    NSString *value = [NSString stringWithFormat:@"{%@}", [items componentsJoinedByString:@","]];
    return value;
}
@end